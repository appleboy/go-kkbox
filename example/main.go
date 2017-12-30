package main

import (
	"fmt"
	"log"
	"os"

	"github.com/appleboy/go-kkbox"
	"github.com/davecgh/go-spew/spew"
)

var clientID = os.Getenv("CLIENT_ID")
var clientSecret = os.Getenv("CLIENT_SECRET")

func main() {
	if clientID == "" || clientSecret == "" {
		log.Fatal("missing client id or secret")
	}
	k, err := kkbox.New(clientID, clientSecret)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("====== kkbox client ======")
	spew.Dump(k)
	fmt.Println("====== kkbox end ======")

	// fetch charts
	charts, err := k.FetchCharts()
	if err != nil {
		fmt.Printf("%#v\n", err)
	}

	fmt.Printf("%#v\n", charts)

	ranks, err := k.FetchChartPlayList("4nUZM-TY2aVxZ2xaA-")
	if err != nil {
		fmt.Printf("%#v\n", err)
	}

	spew.Dump(ranks)

	tracks, err := k.FetchChartPlayListTrack("4nUZM-TY2aVxZ2xaA-", kkbox.Param{
		PerPage: 1,
	})
	if err != nil {
		fmt.Printf("%#v\n", err)
	}

	spew.Dump(tracks)

	// // fetch hits
	// hits, err := k.FetchHits(kkbox.Param{
	// 	PerPage: 2,
	// })
	// if err != nil {
	// 	fmt.Printf("%#v\n", err)
	// }

	// fmt.Println("hit length:", len(hits.Data))
	// // spew.Dump(hits)

	// list, err := k.FetchHitPlayList("DZrC8m29ciOFY2JAm3", kkbox.Param{
	// 	Page:    2,
	// 	PerPage: 2,
	// })
	// if err != nil {
	// 	fmt.Printf("%#v\n", err)
	// }

	// fmt.Println("list length:", len(list.Tracks.Data))
	// // spew.Dump(list)

	// tracks, err := k.FetchHitPlayListTrack("DZrC8m29ciOFY2JAm3", kkbox.Param{
	// 	PerPage: 3,
	// })
	// if err != nil {
	// 	fmt.Printf("%#v\n", err)
	// }

	// fmt.Println("track length:", len(tracks.Data))
	// // spew.Dump(tracks)

	// results, err := k.FetchSearchData("五月天", kkbox.Param{
	// 	PerPage: 1,
	// })
	// if err != nil {
	// 	fmt.Printf("%#v\n", err)
	// }

	// fmt.Println("artist length:", len(results.Artists.Data))
	// fmt.Println("track length:", len(results.Tracks.Data))
	// // spew.Dump(results)

	// track, err := k.FetchTrack("4kxvr3wPWkaL9_y3o_")
	// if err != nil {
	// 	fmt.Printf("%#v\n", err)
	// }
	// spew.Dump(track)

	// album, err := k.FetchAlbum("WpTPGzNLeutVFHcFq6")
	// if err != nil {
	// 	fmt.Printf("%#v\n", err)
	// }
	// spew.Dump(album)

	// albumTrack, err := k.FetchAlbumTrack("KmRKnW5qmUrTnGRuxF", kkbox.Param{
	// 	Page:    1,
	// 	PerPage: 1,
	// })
	// if err != nil {
	// 	fmt.Printf("%#v\n", err)
	// }
	// spew.Dump(albumTrack)
	// fmt.Println("album Track length:", len(albumTrack.Data))

	// artist, err := k.FetchArtist("8q3_xzjl89Yakn_7GB")
	// if err != nil {
	// 	fmt.Printf("%#v\n", err)
	// }
	// spew.Dump(artist)

	// artistAlbums, err := k.FetchArtistAlbum("Cnv_K6i5Ft4y41SxLy", kkbox.Param{
	// 	Page:    2,
	// 	PerPage: 1,
	// })
	// if err != nil {
	// 	fmt.Printf("%#v\n", err)
	// }
	// spew.Dump(artistAlbums)

	// artistTopTracks, err := k.FetchArtistTopTrack("Cnv_K6i5Ft4y41SxLy", kkbox.Param{
	// 	Page:    2,
	// 	PerPage: 1,
	// })
	// if err != nil {
	// 	fmt.Printf("%#v\n", err)
	// }
	// spew.Dump(artistTopTracks)
	// fmt.Println("track length:", len(artistTopTracks.Data))

	// artistRelated, err := k.FetchArtistRelated("8q3_xzjl89Yakn_7GB", kkbox.Param{
	// 	Page:    1,
	// 	PerPage: 1,
	// })
	// if err != nil {
	// 	fmt.Printf("%#v\n", err)
	// }
	// spew.Dump(artistRelated)
	// fmt.Println("artist releated length:", len(artistRelated.Data))

	// sharedPlayList, err := k.FetchSharedPlayList("4nUZM-TY2aVxZ2xaA-")
	// if err != nil {
	// 	fmt.Printf("%#v\n", err)
	// }
	// spew.Dump(sharedPlayList)
	// fmt.Println("artist releated length:", len(sharedPlayList.Tracks.Data))

	// sharedPlayListTrack, err := k.FetchSharedPlayListTrack("4nUZM-TY2aVxZ2xaA-", kkbox.Param{
	// 	PerPage: 1,
	// 	Page:    5,
	// })
	// if err != nil {
	// 	fmt.Printf("%#v\n", err)
	// }
	// spew.Dump(sharedPlayListTrack)
	// fmt.Println("artist releated length:", len(sharedPlayListTrack.Data))

	// featured, err := k.FetchFeatured(kkbox.Param{
	// 	PerPage: 2,
	// })
	// if err != nil {
	// 	fmt.Printf("%#v\n", err)
	// }
	// spew.Dump(featured)
	// fmt.Println("feature play list length:", len(featured.Data))

	// featuredPlayList, err := k.FetchFeaturedPlayList("Wt95My35CqR9hB_FW1")
	// if err != nil {
	// 	fmt.Printf("%#v\n", err)
	// }
	// spew.Dump(featuredPlayList)
	// fmt.Println("feature play list length:", len(featuredPlayList.Tracks.Data))

	// featuredPlayListTrack, err := k.FetchFeaturedPlayListTrack("Wt95My35CqR9hB_FW1", kkbox.Param{
	// 	PerPage: 2,
	// 	Page:    3,
	// })
	// if err != nil {
	// 	fmt.Printf("%#v\n", err)
	// }
	// spew.Dump(featuredPlayListTrack)
	// fmt.Println("feature play list track length:", len(featuredPlayListTrack.Data))

	// featuredCategory, err := k.FetchFeatured(kkbox.Param{
	// 	PerPage: 2,
	// })
	// if err != nil {
	// 	fmt.Printf("%#v\n", err)
	// }
	// spew.Dump(featuredCategory)
	// fmt.Println("feature category length:", len(featuredCategory.Data))

	// featuredSingleCategory, err := k.FetchSingleFeaturedCategory("LXUR688EBKRRZydAWb")
	// if err != nil {
	// 	fmt.Printf("%#v\n", err)
	// }
	// spew.Dump(featuredSingleCategory)
	// fmt.Println("feature category play list length:", len(featuredSingleCategory.Playlists.Data))

	// featuredCategoryPlayList, err := k.FetchFeaturedCategoryPlayList("LXUR688EBKRRZydAWb", kkbox.Param{
	// 	PerPage: 2,
	// })
	// if err != nil {
	// 	fmt.Printf("%#v\n", err)
	// }
	// spew.Dump(featuredCategoryPlayList)
	// fmt.Println("feature category play list length:", len(featuredCategoryPlayList.Data))

	// moodStations, err := k.FetchMoodStationList(kkbox.Param{
	// 	PerPage: 2,
	// })
	// if err != nil {
	// 	fmt.Printf("%#v\n", err)
	// }
	// spew.Dump(moodStations)
	// fmt.Println("mood station list length:", len(moodStations.Data))

	// moodStation, err := k.FetchMoodStation("StGZp2ToWq92diPHS7")
	// if err != nil {
	// 	fmt.Printf("%#v\n", err)
	// }
	// spew.Dump(moodStation)
	// fmt.Println("mood station list length:", len(moodStation.Tracks.Data))

	// genreStations, err := k.FetchGenreStationList(kkbox.Param{
	// 	PerPage: 2,
	// })
	// if err != nil {
	// 	fmt.Printf("%#v\n", err)
	// }
	// spew.Dump(genreStations)
	// fmt.Println("mood station list length:", len(genreStations.Data))

	// genreStation, err := k.FetchGenreStation("StGZp2ToWq92diPHS7")
	// if err != nil {
	// 	fmt.Printf("%#v\n", err)
	// }
	// spew.Dump(genreStation)
	// fmt.Println("mood station list length:", len(genreStation.Tracks.Data))

	releaseCategoryList, err := k.FetchReleaseCategory(kkbox.Param{
		PerPage: 2,
	})
	if err != nil {
		fmt.Printf("%#v\n", err)
	}
	spew.Dump(releaseCategoryList)
	fmt.Println("release Category List length:", len(releaseCategoryList.Data))

	releaseCategory, err := k.FetchSingleReleaseCategory("KrdH2LdyUKS8z2aoxX", kkbox.Param{
		PerPage: 2,
	})
	if err != nil {
		fmt.Printf("%#v\n", err)
	}
	spew.Dump(releaseCategory)
	fmt.Println("release Category List length:", len(releaseCategory.Albums.Data))

	releaseCategoryAlbum, err := k.FetchReleaseCategoryAlbum("KrdH2LdyUKS8z2aoxX", kkbox.Param{
		PerPage: 3,
	})
	if err != nil {
		fmt.Printf("%#v\n", err)
	}
	spew.Dump(releaseCategoryAlbum)
	fmt.Println("release Category List length:", len(releaseCategoryAlbum.Data))
}
