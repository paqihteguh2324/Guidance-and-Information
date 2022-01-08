import React from 'react';
import {
  ImageBackground,
  ScrollView,
  StyleSheet,
  Text,
  View,
} from 'react-native';
import {Dimensions} from 'react-native';
import GambarKotakHeader from '../../assets';
import {
  IconKeranjangdanPesan,
  HeaderTabs,
  ListRidingTips,
  VideoRiding,
} from '../../component';
// import { createAppContainer } from 'react-navigation';
// import { createStackNavigator } from 'react-navigation-stack';
// import { VideoRiding } from './VideoRiding'
// import { RidingTips } from './RidingTips'

const RidingTips = () => {
  return (
    <View style={styles.page}>
      <ImageBackground source={GambarKotakHeader} style={styles.header}>
        <View style={styles.buttonIcon}>
          <IconKeranjangdanPesan title="Keranjang" />
          <IconKeranjangdanPesan title="Pesan" />
        </View>
        <Text style={styles.titleHeader}>Safety</Text>
        {/* <HeaderNavigation/> */}
        <HeaderTabs />
      </ImageBackground>

      {/* INI HALAMAN PDF */}
      {/* <ScrollView showsVerticalScrollIndicator={false}> */}
      <ListRidingTips />
      {/* </ScrollView> */}

      {/* INI HALAMAN VIDEO */}
      {/* <VideoRiding/> */}
    </View>
  );
};

export default RidingTips;

const windowWidth = Dimensions.get('window').width;
const windowHeight = Dimensions.get('window').height;

// const page = createStackNavigator(
//     {
//         RidingTips:{
//             screen: Safety,
//         },
//         VideoRiding:{
//             screen: VideoRiding,
//         },
//     },
//     {
//         headerMode: 'none',
//         initialRouteName: 'RidingTips',
//     }
// );

const styles = StyleSheet.create({
  page: {
    flex: 1,
  },
  header: {
    backgroundColor: '#FFFFFF',
    width: windowWidth,
    height: windowHeight * 0.14,
    shadowColor: '#000',
    elevation: 7,
  },
  titleHeader: {
    //backgroundColor:'red',
    //width: 55,
    color: 'black',
    fontSize: 20,
    fontWeight: 'bold',
    marginTop: -30,
    marginLeft: 20,
  },
  buttonIcon: {
    flexDirection: 'row',
    alignSelf: 'flex-end',
  },
});
