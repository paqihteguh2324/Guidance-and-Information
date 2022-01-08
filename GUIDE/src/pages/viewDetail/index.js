import {StyleSheet, Text, View, Dimensions} from 'react-native';
import React, {Component} from 'react';
import {WebView} from 'react-native-webview';

export default function viewDetail({route}) {
  console.log(route.params.file);
  return (
    <WebView
      source={{
        uri: route.params.file,
      }}
      style={styles.page}
    />
  );
}

const windowWidth = Dimensions.get('window').width;
const windowHeight = Dimensions.get('window').height;

const styles = StyleSheet.create({
  page: {
    flex: 1,
  },
  header: {
    backgroundColor: '#FFFFFF',
    width: windowWidth,
    height: windowHeight * 0.14,
    shadowColor: '#000',
    elevation: 7,
  },
  titleHeader: {
    //backgroundColor:'red',
    //width: 55,
    color: 'black',
    fontSize: 20,
    fontWeight: 'bold',
    marginTop: -30,
    marginLeft: 20,
  },
  buttonIcon: {
    flexDirection: 'row',
    alignSelf: 'flex-end',
  },
});
