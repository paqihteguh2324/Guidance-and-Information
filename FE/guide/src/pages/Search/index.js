import React from 'react'
import { ScrollView, StyleSheet, Text, View } from 'react-native'
import YoutubePlayer from 'react-native-youtube-iframe'

const Search = () => {
    return (
        <ScrollView style={styles.page}>
            
            
            <View>
                <YoutubePlayer
                    height={200}
                    videoId={'Sgbk4YHFTwQ'}
                />
                <Text>Source</Text>
                <Text style={styles.titleVideo}>Wondering what you should pack for your
                        Bicycle Ride Across Georgia? We've compiled a (not exhaustive)
                </Text>

                <YoutubePlayer
                    height={200}
                    videoId={'L0tCJv2s8Jg'}
                />
                <Text>Source</Text>
                <Text style={styles.titleVideo}>A bike trip packing list for your next vacation
                </Text>

                <YoutubePlayer
                    height={200}
                    videoId={'ODmB9LyYzKM'}
                />
                <Text>Source</Text>
                <Text style={styles.titleVideo}>Tips Group cycling etiquette
                </Text>
                
                <YoutubePlayer
                    height={200}
                    videoId={'prYMM7D2qF8'}
                />
                <Text>Source</Text>
                <Text style={styles.titleVideo}>BICYCLE HAND SIGNALS
                </Text>

                <YoutubePlayer
                    height={200}
                    videoId={'6uaEj54flNg'}
                />
                <Text>Source</Text>
                <Text style={styles.titleVideo}>Safety bike riding tips
                </Text>

                <YoutubePlayer
                    height={200}
                    videoId={'EBKeNOBwaVE'}
                />
                <Text>Source</Text>
                <Text style={styles.titleVideo}>how to check your bike condition
                </Text>

            </View>
                       
        </ScrollView>

    )
}

export default Search

const styles = StyleSheet.create({
    page: {
        backgroundColor: '#FFFFFF',
        paddingTop:20,
        paddingLeft:6,
        paddingRight:6,
        marginHorizontal: 5
    },
    titleVideo: {
        flexWrap: 'wrap',
        marginTop:8,
        marginBottom: 40,
        marginRight: 8,
        textAlign: 'justify',
        fontSize: 17,
        color: 'black'

    }
})
