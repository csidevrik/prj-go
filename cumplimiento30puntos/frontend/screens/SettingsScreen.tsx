import * as DocumentPicker from 'expo-document-picker';
import React, { useState } from 'react';
import { View, Button, Text, StyleSheet } from 'react-native';

const SettingsScreen = () => {
  const [selectedPath, setSelectedPath] = useState<string | null>(null);

  const pickFolder = async () => {
    const result = await DocumentPicker.getDocumentAsync({
        type: '*/*',
        multiple: false,
        copyToCacheDirectory: false,
      });

      if (!result.canceled && result.assets && result.assets.length > 0) {
        const uri = result.assets[0].uri;
        const pathParts = uri.split('/');
        pathParts.pop(); // quitamos el nombre del archivo
        const folderUri = pathParts.join('/');
        setSelectedPath(folderUri);
        console.log('📁 Carpeta simulada:', folderUri);
      }
  };


  return (
    <View style={styles.container}>
      <Button title="Seleccionar carpeta raíz" onPress={pickFolder} />
      {selectedPath && (
        <Text style={styles.pathText}>
          Carpeta seleccionada:
          {'\n'}
          {selectedPath}
        </Text>
      )}
    </View>
  );
};

export default SettingsScreen;

const styles = StyleSheet.create({
  container: {
    padding: 20,
  },
  pathText: {
    marginTop: 15,
    backgroundColor: '#eee',
    padding: 10,
  },
});
