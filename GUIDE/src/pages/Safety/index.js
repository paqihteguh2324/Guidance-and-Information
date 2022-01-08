import React from 'react';
import {StyleSheet} from 'react-native';
import {createNativeStackNavigator} from '@react-navigation/native-stack';
import RidingTips from '../RidingTips';
import VideoRiding from '../VideoRiding';
import viewDetail from '../viewDetail';

const RootStack = createNativeStackNavigator();

const Safety = () => {
  return (
    <RootStack.Navigator style={{flex: 1}}>
      <RootStack.Screen
        name="Riding Tips"
        component={RidingTips}
        options={{headerShown: false}}
      />
      <RootStack.Screen
        name="Video Riding"
        component={VideoRiding}
        options={{headerShown: false}}
      />
      <RootStack.Screen
        name="View Detail"
        component={viewDetail}
        options={{headerShown: true}}
      />
    </RootStack.Navigator>
  );
};

export default Safety;
