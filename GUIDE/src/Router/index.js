import React from 'react'
import { StyleSheet, Text, View } from 'react-native'
import { createNativeStackNavigator } from '@react-navigation/native-stack';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';

import Feed from '../pages/Feed';
import Search from '../pages/Search';
import Home from '../pages/Home';
import Safety from '../pages/Safety';
import Profile from '../pages/Profile';
import BottomNavigation from '../component/BottomNavigation';
import VideoRiding from '../pages/VideoRiding';

const Stack = createNativeStackNavigator();
const Tab = createBottomTabNavigator();

const MainApp = () => {
    return (
        <Tab.Navigator tabBar={props => <BottomNavigation {...props} />}>
            <Tab.Screen name="Feed" component={Feed} options={{ headerShown: false }} />
            <Tab.Screen name="Search" component={Search} options={{ headerShown: false }} />
            <Tab.Screen name=" " component={Home} options={{ headerShown: false }} />
            <Tab.Screen name="Safety" component={Safety} options={{ headerShown: false }} />
            <Tab.Screen name="Profile" component={Profile} options={{ headerShown: false }} />
        </Tab.Navigator>
    )
}

const Router = () => {
    return (
        <Stack.Navigator>
            <Stack.Screen name="MainApp" component={MainApp} options={{ headerShown: false }} />
            <Stack.Screen name="VideoRiding" component={VideoRiding} />
        </Stack.Navigator>

    )

}

export default Router
