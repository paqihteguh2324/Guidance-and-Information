import React, {useState, useEffect} from 'react';
import {
  View,
  Text,
  TouchableOpacity,
  Image,
  Alert,
  ActivityIndicator,
  FlatList,
  Dimensions,
} from 'react-native';
import {useNavigation} from '@react-navigation/native';

function ListRidingTips() {
  const [isLoading, setLoading] = useState(true);
  const [data, setData] = useState([]);
  const getMovies = async () => {
    try {
      const response = await fetch('http://192.168.43.217:8080/api/artikel');
      const json = await response.json();
      setData(json.data);
    } catch (error) {
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    getMovies();
  }, []);

  const DataRidingTips = [{}];
  return (
    <View>
      {isLoading ? (
        <ActivityIndicator />
      ) : (
        <FlatList
          showsVerticalScrollIndicator={false}
          style={{marginBottom: 100}}
          data={data}
          keyExtractor={({artikel_id2}, index) => artikel_id2}
          renderItem={({item}) => (
            <View>
              <TouchableOpacity
                activeOpacity={1}
                style={{
                  marginBottom: 9,
                  paddingHorizontal: 17,
                  flex: 1,
                }}>
                {DataRidingTips.map((RidingTips, index) => (
                  <View
                    key={index}
                    style={{
                      marginTop: 10,
                      padding: 15,
                      backgroundColor: '#FFF',
                      flexDirection: 'row',
                    }}>
                    <PDFImage image={{uri: item.artikel_image}} />
                    <PDFDescription name={item.artikel_headings} />
                    <View>
                      <ViewDetail
                        text="View Detail"
                        btnColor="#0C8EFF"
                        textColor="#FFF"
                        textViewDetail={item.artikel_desc}
                      />
                    </View>
                  </View>
                ))}
              </TouchableOpacity>
            </View>
          )}
        />
      )}
    </View>
  );
}

const PDFImage = props => (
  <>
    <Image source={props.image} style={{width: 175, height: 175}} />
  </>
);

const PDFDescription = props => (
  <View>
    <Text
      style={{
        color: 'black',
        width: 190,
        padding: 10,
        textAlign: 'left',
        fontSize: 17,
      }}>
      {props.name}
    </Text>
  </View>
);

const ViewDetail = props => {
  const navigation = useNavigation();
  const [activeTab, setActiveTab] = useState('view Detail');
  return (
    <TouchableOpacity
      style={{
        backgroundColor: props.btnColor,
        paddingHorizontal: 38,
        paddingVertical: 7,
        borderRadius: 20,
        width: 173,
        height: 34,
        //alignSelf: "flex-start",
        marginTop: 140,
        marginLeft: -180,
      }}
      // onPress={() => Alert.alert(props.textViewDetail)}
      onPress={() =>
        navigation.navigate('View Detail', {file: props.textViewDetail})
      }>
      <Text style={{color: props.textColor, alignSelf: 'center'}}>
        {props.text}
      </Text>
    </TouchableOpacity>
  );
};
export default ListRidingTips;
