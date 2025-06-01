import React from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createDrawerNavigator } from '@react-navigation/drawer';
import DashboardScreen from './screens/HomeScreen';
import SettingsScreen from './screens/SettingsScreen';
import { Text, View } from 'react-native';

const Drawer = createDrawerNavigator();

export default function App() {
  return (
    <NavigationContainer>
      <Drawer.Navigator
        screenOptions={{
          headerStyle: { backgroundColor: '#eee' },
          headerTitle: 'CUMPLIMIENTO30PUNTOS',
          headerTitleAlign: 'center',
        }}
      >
        <Drawer.Screen name="Dashboard" component={DashboardScreen} />
        <Drawer.Screen name="Configuración" component={SettingsScreen} />
      </Drawer.Navigator>
    </NavigationContainer>
  );
}
