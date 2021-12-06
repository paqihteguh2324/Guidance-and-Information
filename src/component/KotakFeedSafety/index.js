import React from 'react'
import { ScrollView, StyleSheet, Text, Button, View, Alert, ImageBackground } from 'react-native'
 import { ButtonViewDetail } from '../../component'
import { fotoFeed1, fotoFeed2, fotoFeed3, fotoFeed4 } from '../../assets'
import ListRidingTips from '../ListRidingTips'
const KotakFeedSafety = () => {
    return (
        <ScrollView showsVerticalScrollIndicator={false} style={styles.page}>
            
            <ListRidingTips/>
            {/* <View style={styles.kotakFeed}>
                <ImageBackground source={fotoFeed1} style={styles.fotoKonten}>

                </ImageBackground>
                <Text style={styles.titleFeed}>6 Safety points,  for {"\n"}
                    children while out riding {"\n"}their bikes.
                </Text>
                <View style={styles.buttonView}>
                <View>
                    <ButtonViewDetail/>
                </View>
            </View>
                
            </View>
            
            
            <View style={styles.kotakFeed}>
                <ImageBackground source={fotoFeed2} style={styles.fotoKonten}>

                </ImageBackground>
                <Text style={styles.titleFeed}>7 Gesture bicycle hand {"\n"}
                    signals</Text>
            </View>
            

            <View style={styles.kotakFeed}>
                <ImageBackground source={fotoFeed3} style={styles.fotoKonten}>

                </ImageBackground>
                <Text style={styles.titleFeed}>Tips what to wear on {"\n"}the
                    bike</Text>
            </View>
            <View style={styles.kotakFeed}>
                <ImageBackground source={fotoFeed4} style={styles.fotoKonten}>

                </ImageBackground>
                <Text style={styles.titleFeed}>Tips Group cycling {"\n"}
                    etiquette</Text>
            </View> */}
            
        </ScrollView>

    )
}

export default KotakFeedSafety

const styles = StyleSheet.create({
    // page: {
    //     //marginTop: -0,
    //     paddingTop:20,
    //     width: 393,
    //     marginHorizontal: 5
    // },
    // kotakFeed: {
    //     flexWrap: 'wrap',
    //     backgroundColor: '#FFFFFF',
    //     height: 197,
    //     width: 393,
    //     borderRadius: 5,
    //     marginVertical: 9,
    //     alignItems: "center",
    //     shadowColor: "#000",
    //     elevation: 5,
    // },
    // fotoKonten: {
    //     //backgroundColor: '#C4C4C4',
    //     height: 175,
    //     width: 175,
    //     marginLeft: 13,
    //     marginTop: 9,
    //     marginBottom: 13

    // },
    // titleFeed: {
    //     flexWrap: 'wrap',
    //     marginTop: 31,
    //     marginBottom: 97,
    //     marginRight: 8,
    //     textAlign: 'justify',
    //     fontSize: 17,
    //     color: 'black'

    // }
    
})
