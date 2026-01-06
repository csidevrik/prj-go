use clap::{Parser, Subcommand};
use colored::*;
use dialoguer::{theme::ColorfulTheme, Select, Input};
use serde::{Deserialize, Serialize};
use std::fs;
use std::path::PathBuf;
use std::process::Command;

#[derive(Parser)]
#[command(name = "palo-launcher")]
#[command(about = "Herramienta CLI para gestionar conexiones al Firewall Palo Alto", long_about = None)]
struct Cli {
    #[command(subcommand)]
    command: Option<Commands>,
}

#[derive(Subcommand)]
enum Commands {
    /// Configurar la IP del firewall y navegadores
    Config,
    /// Conectar al firewall (modo interactivo)
    Connect,
    /// Mostrar configuración actual
    Show,
}

#[derive(Serialize, Deserialize, Clone)]
struct Config {
    firewall_ip: String,
    firefox_profile: String,
    brave_profile: String,
    chrome_profile: String,
    default_browsers: Vec<String>,
}

impl Default for Config {
    fn default() -> Self {
        Config {
            firewall_ip: String::new(),
            firefox_profile: "palo-firefox".to_string(),
            brave_profile: "palo-brave".to_string(),
            chrome_profile: "palo-chrome".to_string(),
            default_browsers: vec!["firefox".to_string(), "brave".to_string()],
        }
    }
}

struct PaloLauncher {
    config_path: PathBuf,
    config: Config,
}

impl PaloLauncher {
    fn new() -> Result<Self, Box<dyn std::error::Error>> {
        let config_dir = dirs::config_dir()
            .ok_or("No se pudo obtener el directorio de configuración")?
            .join("palo-launcher");
        
        fs::create_dir_all(&config_dir)?;
        let config_path = config_dir.join("config.json");
        
        let config = if config_path.exists() {
            let content = fs::read_to_string(&config_path)?;
            serde_json::from_str(&content)?
        } else {
            Config::default()
        };
        
        Ok(PaloLauncher { config_path, config })
    }
    
    fn save_config(&self) -> Result<(), Box<dyn std::error::Error>> {
        let json = serde_json::to_string_pretty(&self.config)?;
        fs::write(&self.config_path, json)?;
        Ok(())
    }
    
    fn configure(&mut self) -> Result<(), Box<dyn std::error::Error>> {
        println!("{}", "=== Configuración de Palo Alto Launcher ===".green().bold());
        println!();
        
        let ip: String = Input::with_theme(&ColorfulTheme::default())
            .with_prompt("IP de administración del Palo Alto")
            .default(self.config.firewall_ip.clone())
            .interact_text()?;
        
        self.config.firewall_ip = ip;
        
        println!("\n{}", "Selecciona los navegadores a usar (por defecto):".cyan());
        let browsers = vec!["Firefox + Brave", "Firefox + Chrome", "Solo Firefox (2 ventanas)"];
        let selection = Select::with_theme(&ColorfulTheme::default())
            .items(&browsers)
            .default(0)
            .interact()?;
        
        self.config.default_browsers = match selection {
            0 => vec!["firefox".to_string(), "brave".to_string()],
            1 => vec!["firefox".to_string(), "chrome".to_string()],
            2 => vec!["firefox".to_string(), "firefox2".to_string()],
            _ => vec!["firefox".to_string(), "brave".to_string()],
        };
        
        self.save_config()?;
        println!("\n{}", "✓ Configuración guardada exitosamente".green());
        
        Ok(())
    }
    
    fn show_config(&self) {
        println!("{}", "=== Configuración Actual ===".cyan().bold());
        println!("IP del Firewall: {}", self.config.firewall_ip.yellow());
        println!("Navegadores: {}", self.config.default_browsers.join(" + ").yellow());
        println!("Perfil Firefox: {}", self.config.firefox_profile.yellow());
        println!("Perfil Brave: {}", self.config.brave_profile.yellow());
        println!("Perfil Chrome: {}", self.config.chrome_profile.yellow());
    }
    
    fn interactive_connect(&self) -> Result<(), Box<dyn std::error::Error>> {
        if self.config.firewall_ip.is_empty() {
            println!("{}", "⚠ Primero debes configurar la IP del firewall".red());
            println!("Ejecuta: {} config", "palo-launcher".cyan());
            return Ok(());
        }
        
        println!("{}", "=== Conexión a Palo Alto ===".green().bold());
        println!("Firewall: {}", self.config.firewall_ip.yellow());
        println!();
        
        let sections = vec![
            "Dashboard Principal",
            "Monitor",
            "Policies",
            "Objects",
            "Network",
            "Logs de Tráfico",
            "Configuración Personalizada",
        ];
        
        let selection = Select::with_theme(&ColorfulTheme::default())
            .with_prompt("Selecciona la sección a abrir")
            .items(&sections)
            .default(0)
            .interact()?;
        
        let url = match selection {
            0 => format!("https://{}", self.config.firewall_ip),
            1 => format!("https://{}/#monitor", self.config.firewall_ip),
            2 => format!("https://{}/#policies", self.config.firewall_ip),
            3 => format!("https://{}/#objects", self.config.firewall_ip),
            4 => format!("https://{}/#network", self.config.firewall_ip),
            5 => format!("https://{}/#monitor/traffic", self.config.firewall_ip),
            6 => {
                let custom: String = Input::with_theme(&ColorfulTheme::default())
                    .with_prompt("Ingresa la ruta (ej: #monitor/system)")
                    .interact_text()?;
                format!("https://{}/{}", self.config.firewall_ip, custom)
            }
            _ => format!("https://{}", self.config.firewall_ip),
        };
        
        self.open_browsers(&url)?;
        
        Ok(())
    }
    
    fn open_browsers(&self, url: &str) -> Result<(), Box<dyn std::error::Error>> {
        println!("\n{}", "Abriendo navegadores...".cyan());
        
        for (i, browser) in self.config.default_browsers.iter().enumerate() {
            match browser.as_str() {
                "firefox" | "firefox2" => {
                    let profile = if browser == "firefox2" {
                        format!("{}-2", self.config.firefox_profile)
                    } else {
                        self.config.firefox_profile.clone()
                    };
                    self.open_firefox(url, &profile)?;
                    println!("  {} Firefox (perfil: {})", "✓".green(), profile);
                }
                "brave" => {
                    self.open_brave(url, &self.config.brave_profile)?;
                    println!("  {} Brave (perfil: {})", "✓".green(), self.config.brave_profile);
                }
                "chrome" => {
                    self.open_chrome(url, &self.config.chrome_profile)?;
                    println!("  {} Chrome (perfil: {})", "✓".green(), self.config.chrome_profile);
                }
                _ => {}
            }
            
            // Pequeña pausa entre navegadores
            if i < self.config.default_browsers.len() - 1 {
                std::thread::sleep(std::time::Duration::from_millis(500));
            }
        }
        
        println!("\n{}", "¡Navegadores abiertos exitosamente!".green().bold());
        println!("{}", "Las sesiones se mantendrán en los perfiles específicos.".yellow());
        
        Ok(())
    }
    
    fn open_firefox(&self, url: &str, profile: &str) -> Result<(), Box<dyn std::error::Error>> {
        // Intenta con 'firefox' primero, luego 'firefox-esr'
        let firefox_cmd = if Command::new("firefox").arg("--version").output().is_ok() {
            "firefox"
        } else {
            "firefox-esr"
        };
        
        Command::new(firefox_cmd)
            .arg("-P")
            .arg(profile)
            .arg("--new-instance")
            .arg(url)
            .spawn()?;
        
        Ok(())
    }
    
    fn open_brave(&self, url: &str, profile: &str) -> Result<(), Box<dyn std::error::Error>> {
        // Rutas comunes de Brave en Linux
        let brave_paths = vec![
            "brave-browser",
            "brave",
            "/usr/bin/brave-browser",
            "/usr/bin/brave",
        ];
        
        for brave_cmd in brave_paths {
            if Command::new(brave_cmd).arg("--version").output().is_ok() {
                let profile_dir = dirs::home_dir()
                    .ok_or("No se pudo obtener el directorio home")?
                    .join(".config")
                    .join("BraveSoftware")
                    .join("Brave-Browser")
                    .join(profile);
                
                Command::new(brave_cmd)
                    .arg(format!("--user-data-dir={}", profile_dir.display()))
                    .arg(url)
                    .spawn()?;
                
                return Ok(());
            }
        }
        
        Err("Brave no está instalado o no se encuentra en el PATH".into())
    }
    
    fn open_chrome(&self, url: &str, profile: &str) -> Result<(), Box<dyn std::error::Error>> {
        // Rutas comunes de Chrome en Linux
        let chrome_paths = vec![
            "google-chrome",
            "google-chrome-stable",
            "chromium",
            "chromium-browser",
        ];
        
        for chrome_cmd in chrome_paths {
            if Command::new(chrome_cmd).arg("--version").output().is_ok() {
                let profile_dir = dirs::home_dir()
                    .ok_or("No se pudo obtener el directorio home")?
                    .join(".config")
                    .join("google-chrome")
                    .join(profile);
                
                Command::new(chrome_cmd)
                    .arg(format!("--user-data-dir={}", profile_dir.display()))
                    .arg(url)
                    .spawn()?;
                
                return Ok(());
            }
        }
        
        Err("Chrome/Chromium no está instalado o no se encuentra en el PATH".into())
    }
}

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let cli = Cli::parse();
    let mut launcher = PaloLauncher::new()?;
    
    match cli.command {
        Some(Commands::Config) => {
            launcher.configure()?;
        }
        Some(Commands::Show) => {
            launcher.show_config();
        }
        Some(Commands::Connect) => {
            launcher.interactive_connect()?;
        }
        None => {
            // Modo interactivo por defecto
            launcher.interactive_connect()?;
        }
    }
    
    Ok(())
}
