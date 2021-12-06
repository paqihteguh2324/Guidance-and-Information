import React from 'react'
import { StyleSheet, Text, View,Button,Alert } from 'react-native'

const ButtonViewDetail = () => {
    return (
        <View style={styles.container}>
            <Button
                title="View Detail"
                onPress={() => Alert.alert('buka pdf')}
            />
        </View>
    )
}

export default ButtonViewDetail

const styles = StyleSheet.create({
    container: {
        //flexDirection: 'column',
        //justifyContent: 'space-evenly',
        marginTop:150,
        marginLeft: -250,
        //marginBottom: -50,
        //marginRight: -30,
        alignSelf: 'flex-end',
        width: 173,
        borderRadius: 20
    },
    buttonView2: {
        
        marginTop: 10
    }
})
