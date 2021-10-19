import React from 'react'
import { StyleSheet, Text, View, TouchableOpacity} from 'react-native'
import TabItems from '../TabItems';

const BottomNavigation = ({ state, descriptors, navigation }) => {
    const focusedOptions = descriptors[state.routes[state.index].key].option;

    // if(focusedOptions.tabBarVisible === false){
    //     return null;
    // }
    return (
        <View style={styles.container}>
          {state.routes.map((route, index) => {
            const { options } = descriptors[route.key];
            const label =
              options.tabBarLabel !== undefined
                ? options.tabBarLabel
                : options.title !== undefined
                ? options.title
                : route.name;
    
            const isFocused = state.index === index;
    
            const onPress = () => {
              const event = navigation.emit({
                type: 'tabPress',
                target: route.key,
                canPreventDefault: true,
              });
    
              if (!isFocused && !event.defaultPrevented) {
                // The `merge: true` option makes sure that the params inside the tab screen are preserved
                navigation.navigate({ name: route.name, merge: true });
              }
            };
    
            const onLongPress = () => {
              navigation.emit({
                type: 'tabLongPress',
                target: route.key,
              });
            };
    
            return (
              <TabItems
              key={index}
              label={label}
              isFocused={isFocused}
              onPress={onPress}
              onLongPress={onLongPress}
              />
            );
          })}
        </View>
      );
}

export default BottomNavigation

const styles = StyleSheet.create({
    container: {
        flexDirection: 'row',
        backgroundColor: '#FFFFFF',
        justifyContent: 'space-between',
        paddingHorizontal: 38,
        paddingVertical:15,
        borderTopLeftRadius: 29,
        borderTopRightRadius: 29
        // paddingTop: 26,
        // paddingBottom: 21,
    }
})
