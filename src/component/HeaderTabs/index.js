import React,{useState} from "react"
import { View, Text, TouchableOpacity } from 'react-native'
import { HeaderTabs } from '..'

export default function index() {
    const[activeTab, setActiveTab] = useState("Riding Tips");
    return (
        <View style={{
            marginLeft: 15,
            flexDirection:"row",
            alignSelf: "flex-start"}}>
            <HeaderButton 
            text="Riding Tips" 
            btnColor="#0C8EFF" 
            textColor="#FFF"
            activeTab={activeTab}
            setActiveTab={setActiveTab}/>
            <HeaderButton 
            text="Video Riding" 
            btnColor="#FFF" 
            textColor="#0C8EFF"
            activeTab={activeTab}
            setActiveTab={setActiveTab}/>
        </View>
    )
}

const HeaderButton = (props) => (
    <TouchableOpacity style={{
        backgroundColor: props.activeTab === props.text ? '#0C8EFF' : '#F0F2F5',
        paddingVertical: 5,
        paddingHorizontal: 15,
        marginHorizontal: 3,
        //marginTop: 15,
        marginVertical:10,
        borderRadius: 5}}
        onPress={() => props.setActiveTab(props.text)}
        >

        <Text style={{color: props.activeTab === props.text ? '#FFF' : '#757575'}}>{props.text}</Text>
    </TouchableOpacity>
);
