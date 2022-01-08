import React, {useState, useEffect} from 'react';
import {
  StyleSheet,
  Text,
  View,
  ActivityIndicator,
  FlatList,
} from 'react-native';
import YoutubePlayer from 'react-native-youtube-iframe';

const ListVideoRiding = () => {
  const [isLoading, setLoading] = useState(true);
  const [data, setData] = useState([]);
  const getMovies = async () => {
    try {
      const response = await fetch('http://192.168.43.217:8080/api/video');
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

  const DataVideoRiding = [{}];

  return (
    <View>
      {isLoading ? (
        <ActivityIndicator />
      ) : (
        <FlatList
          showsVerticalScrollIndicator={false}
          style={{marginBottom: 100}}
          data={data}
          keyExtractor={({video_id}, index) => video_id}
          renderItem={({item}) => (
            <View
              style={{
                marginTop: 19,
                padding: 15,
                backgroundColor: '#FFF',
              }}>
              <YoutubePlayer height={230} videoId={item.video_link} />
              <Text> source {item.video_desc}</Text>
              <Text style={styles.titleVideo}>{item.video_headings}</Text>
            </View>
          )}
        />
      )}
    </View>
  );
};

export default ListVideoRiding;

const styles = StyleSheet.create({
  page: {
    backgroundColor: '#FFFFFF',
    paddingTop: 20,
    paddingLeft: 6,
    paddingRight: 6,
    marginHorizontal: 5,
  },
  titleVideo: {
    flexWrap: 'wrap',
    marginLeft: 4,
    marginBottom: 17,
    marginRight: 8,
    textAlign: 'justify',
    fontSize: 17,
    color: 'black',
  },
});
