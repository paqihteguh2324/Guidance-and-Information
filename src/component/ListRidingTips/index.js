import React from 'react'
import { View, Text, TouchableOpacity, Image, Alert} from 'react-native'
import { fotoFeed1 } from '../../assets';

export default function ListRidingTips() {
    return (
        <TouchableOpacity activeOpacity={1} style={{marginBottom:9,paddingHorizontal:17}}>
            <View style={{
                marginTop: 10,
                padding:15,
                backgroundColor:"#FFF",
                flexDirection: "row"}}>
                <PDFImage/>
                <PDFDescription/>
                
            </View>
            <View>
                 <ViewDetail 
                    text="View Detail" 
                    btnColor="#0C8EFF" 
                    textColor="#FFF"
                />
            </View>
        </TouchableOpacity>
        

    );
}

const PDFImage = () => (
    <>
    <Image
        source={fotoFeed1}
        style={{width:175,height:175}}
    />
    
    </>
);

const PDFDescription =()=> (
    <View>
        <Text style={{color: "black",width:190,padding:10,textAlign:"justify",fontSize:17}}>6 Safety points,  for children while out riding their bikes.</Text>
    </View>
);

const ViewDetail = (props) => (
    <TouchableOpacity style={{
        backgroundColor: props.btnColor,
        paddingHorizontal: 38,
        paddingVertical:7,
        borderRadius:20,
        width: 173,
        height:34,
        alignSelf: "flex-end",
        marginTop: -50,
        marginRight: 7
    }}
   
    >
        <Text style={{color:props.textColor,alignSelf:"center"}}>{props.text}</Text>
    </TouchableOpacity>
)