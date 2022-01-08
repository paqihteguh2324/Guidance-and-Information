import React from 'react'
import { StyleSheet, Text, View, ImageBackground, TouchableOpacity,Alert } from 'react-native'
import { IconKeranjang, IconKeranjangAktif, IconPesan, IconPesanAktif } from '../../assets'

const IconKeranjangdanPesan = ({ isFocused, onPress, onLongPress, title }) => {
    const Icon = () => {
        if (title === "Keranjang") return <IconKeranjangAktif />
        if (title === "Pesan") return <IconPesanAktif />

        return <IconKeranjang />
    }
    return (
        <TouchableOpacity
            style={styles.container}>
            <Icon />
            
        </TouchableOpacity>
    )
}

export default IconKeranjangdanPesan

const styles = StyleSheet.create({
    container: {
        flexDirection: 'row-reverse',

        width: 5,
        //backgroundColor: 'red',
        marginTop: 20,
        marginHorizontal: 18,
        //justifyContent: 'space-between'
    },
    keranjangDanPesan: {
        marginRight: 21
    }
})
