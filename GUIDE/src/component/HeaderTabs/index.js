import {NavigationContainer, useNavigation} from '@react-navigation/native';
import React, {useState} from 'react';
import {View, Text, TouchableOpacity} from 'react-native';
import {RidingTips, VideoRiding, Safety} from '../../pages/Safety';
import {HeaderTabs} from '..';

export default function index() {
  const [activeTab, setActiveTab] = useState('Riding Tips');

  return (
    <View
      style={{
        marginLeft: 15,
        flexDirection: 'row',
        alignSelf: 'flex-start',
      }}>
      <HeaderButton
        text="Riding Tips"
        btnColor="#0C8EFF"
        textColor="#FFF"
        nav="Riding Tips"
        activeTab={activeTab}
        setActiveTab={setActiveTab}
      />
      <HeaderButton
        text="Video Riding"
        btnColor="#FFF"
        textColor="#0C8EFF"
        nav="Video Riding"
        activeTab={activeTab}
        setActiveTab={setActiveTab}
      />
    </View>
  );
}

const HeaderButton = props => {
  const navigation = useNavigation();
  return (
    <TouchableOpacity
      style={{
        backgroundColor: props.activeTab === props.text ? '#F0F2F5' : '#0C8EFF',
        paddingVertical: 5,
        paddingHorizontal: 15,
        marginHorizontal: 3,
        //marginTop: 15,
        marginVertical: 10,
        borderRadius: 5,
      }}
      onPress={() =>
        navigation.navigate(props.nav) || props.setActiveTab(props.text)
      }>
      <Text
        style={{color: props.activeTab === props.text ? '#757575' : '#FFF'}}>
        {props.text}
      </Text>
    </TouchableOpacity>
  );
};
