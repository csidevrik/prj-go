import { StatusBar } from 'expo-status-bar';
import { StyleSheet, Text, View } from 'react-native';
import { SafeAreaView } from 'react-native';
import DashboardTable from './components/DashboardTable';

export default function App() {
  return (
    // <View style={styles.container}>
    //   <Text>Open up App.tsx to start working on your app!</Text>
    //   <StatusBar style="auto" />
      
    // </View>
    <SafeAreaView>
      <DashboardTable />
    </SafeAreaView>

  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },
});
