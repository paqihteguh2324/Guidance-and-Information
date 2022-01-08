import React from 'react'
import { StyleSheet, Text, View, Dimensions } from 'react-native'

const HeaderNavigation = () => {
    return (
        <View style={styles.container}>
            <View style={styles.headerNav}>
                <Text style={styles.titleNav}>Riding Tips</Text>
            </View>
            
            <View style={styles.headerNav2}>
                <Text style={styles.titleNav2}>Video Riding</Text>
            </View>
            
            <View style={styles.headerNav2}>
                <Text style={styles.titleNav2}>e-Sertifikasi</Text>
            </View>
        </View>
    )
}

export default HeaderNavigation

const windowWidth = Dimensions.get('window').width;
const windowHeight = Dimensions.get('window').height;

const styles = StyleSheet.create({
    container: {
        flexDirection: 'row',
        justifyContent: 'space-between',
        marginLeft: 20,
        //marginBottom: 10,
        marginRight: 106,
        
    },
    headerNav: {
        backgroundColor: '#0C8EFF',
        marginRight: 12,
        marginBottom: 10,
        marginTop: 5,
        paddingHorizontal: 13,
        paddingTop:7,
        paddingBottom: 8,
        borderRadius: 4
    },
    headerNav2: {
        backgroundColor: '#F0F2F5',
        marginRight: 12,
        marginBottom: 10,
        marginTop: 5,
        paddingHorizontal: 13,
        paddingTop:7,
        paddingBottom: 8,
        borderRadius: 4
    },
    titleNav: {
        color: '#FFF',
        fontSize: 13.5
    },
    titleNav2: {
        color: '#757575',
        fontSize: 13.5
    }
})
