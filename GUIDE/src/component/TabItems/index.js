import React from 'react'
import { StyleSheet, Text, TouchableOpacity} from 'react-native'

import {IconFeed,IconFeedAktif,IconSearch,IconSearchAktif,IconHome,IconSafety,IconSafetyAktif,
    IconProfile,IconProfileAktif}from '../../assets/icons'


const TabItems = ({isFocused,onPress,onLongPress, label,title}) => {
    const Icons = () => {
        if(label === "Feed") return isFocused ? <IconFeedAktif /> : <IconFeed/>
        if(label === "Search") return isFocused ? <IconSearchAktif /> : <IconSearch/>
         if(label === " ") return isFocused ? <IconHome /> : <IconHome/>
         if(label === "Safety") return isFocused ? <IconSafetyAktif /> : <IconSafety/>
         if(label === "Profile") return isFocused ? <IconProfileAktif /> : <IconProfile/>

        return <IconSafety/>
    }
    return (
        <TouchableOpacity
            onPress={onPress}
            onLongPress={onLongPress}
            style={styles.container}>
            <Icons/>
            <Text style={styles.text(isFocused)}>
                {label}
            </Text>
        </TouchableOpacity>
    )
}

export default TabItems

const styles = StyleSheet.create({
    container: {
        alignItems: 'center'
    },
    text: (isFocused) => ({
        fontSize: 12,
        color: isFocused ? '#0C8EFF' : '#222',
        marginTop: 8
    })
})
