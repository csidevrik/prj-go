import React, { useState } from 'react';
import { View, Text, ScrollView, TouchableOpacity, StyleSheet, Modal } from 'react-native';
import { Card } from 'react-native-paper';

const MESES = ['ENE', 'FEB', 'MAR', 'ABR', 'MAY', 'JUN', 'JUL', 'AGO', 'SEP', 'OCT', 'NOV', 'DIC'];
const PUNTOS = Array.from({ length: 30 }, (_, i) => i + 1);

// Datos simulados
const mockData: Record<number, Record<number, any>> = {
  1: { 0: { estado: 'Sí' }, 1: { estado: 'Sí' } },
  2: { 1: { estado: 'X' } },
  3: { 0: { estado: 'Sí' }, 2: { estado: 'Sa', archivo: 'D1_INF_CTP:002.pdf', pagina: 7 } },
  4: { 3: { estado: 'No' } },
  5: { 0: { estado: 'Sí' } }
};

const DashboardTable = () => {
  const [tooltipData, setTooltipData] = useState<{ archivo: string; pagina: number } | null>(null);
  const [modalVisible, setModalVisible] = useState(false);

  const handleCellPress = (punto: number, mes: number) => {
    const cell = mockData[punto]?.[mes];
    if (cell?.archivo) {
      setTooltipData({ archivo: cell.archivo, pagina: cell.pagina });
      setModalVisible(true);
    }
  };

  return (
    <View style={styles.container}>
      <ScrollView horizontal>
        <View>
          <View style={styles.headerRow}>
            <Text style={[styles.headerCell, styles.fixedCell]}>PUNTO</Text>
            {MESES.map((mes, idx) => (
              <Text key={idx} style={styles.headerCell}>{mes}</Text>
            ))}
          </View>

          <ScrollView style={{ maxHeight: 500 }}>
            {PUNTOS.map((punto) => (
              <View key={punto} style={styles.row}>
                <Text style={[styles.cell, styles.fixedCell]}>{punto}</Text>
                {MESES.map((_, mesIdx) => {
                  const estado = mockData[punto]?.[mesIdx]?.estado || '';
                  return (
                    <TouchableOpacity
                      key={mesIdx}
                      style={styles.cell}
                      onPress={() => handleCellPress(punto, mesIdx)}
                    >
                      <Text>{estado}</Text>
                    </TouchableOpacity>
                  );
                })}
              </View>
            ))}
          </ScrollView>
        </View>
      </ScrollView>

      <Modal
        visible={modalVisible}
        transparent
        animationType="fade"
        onRequestClose={() => setModalVisible(false)}
      >
        <View style={styles.modalOverlay}>
          <Card style={styles.tooltip}>
            <Text>{tooltipData?.archivo}</Text>
            <Text>Página {tooltipData?.pagina}</Text>
          </Card>
        </View>
      </Modal>
    </View>
  );
};

export default DashboardTable;

const styles = StyleSheet.create({
  container: {
    padding: 10,
    backgroundColor: '#fff'
  },
  headerRow: {
    flexDirection: 'row',
    backgroundColor: '#eee'
  },
  row: {
    flexDirection: 'row'
  },
  headerCell: {
    width: 70,
    padding: 6,
    fontWeight: 'bold',
    textAlign: 'center',
    borderWidth: 0.5
  },
  cell: {
    width: 70,
    padding: 6,
    textAlign: 'center',
    borderWidth: 0.5
  },
  fixedCell: {
    backgroundColor: '#ddd'
  },
  modalOverlay: {
    flex: 1,
    backgroundColor: '#000000aa',
    justifyContent: 'center',
    alignItems: 'center'
  },
  tooltip: {
    padding: 16,
    width: 250
  }
});
