import React from 'react'
import { Button, ImageBackground, ScrollView, ScrollViewComponent, StyleSheet, Text, View } from 'react-native'
import { Dimensions } from 'react-native';
import GambarKotakHeader from '../../assets'
import { KotakFeedSafety,IconKeranjangdanPesan} from '../../component';
import HeaderNavigation from '../../component/HeaderNavigation';

const Safety = () => {
    return (
        <View style={styles.page}>
            <ImageBackground source={GambarKotakHeader} style={styles.header}>
                <View style={styles.buttonIcon}>
                    <IconKeranjangdanPesan title="Keranjang"/>
                    <IconKeranjangdanPesan title="Pesan"/>
                </View>
                <Text style={styles.titleHeader}>Safety</Text>
                <HeaderNavigation/>
            </ImageBackground>
            
            
            <KotakFeedSafety />

        </View>
    )
}

export default Safety

const windowWidth = Dimensions.get('window').width;
const windowHeight = Dimensions.get('window').height;

const styles = StyleSheet.create({
    page: {
        flex: 1,
    },
    header: {
        backgroundColor: '#FFFFFF',
        width: windowWidth,
        height: windowHeight * 0.14,
        shadowColor: "#000",
        elevation: 7
    },
    titleHeader: {
        //backgroundColor:'red',
        width: 55,
        color: 'black',
        fontSize: 18,
        fontWeight: 'bold',
        marginTop: -20,
        marginLeft: 20
    },
    buttonIcon: {
        flexDirection:'row',
        alignSelf: 'flex-end',
    }
    
})
