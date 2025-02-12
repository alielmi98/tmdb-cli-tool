package cmd

import (
	"fmt"
	"os"

	"github.com/alielmi98/tmdb-cli-tool/internal/api"
	"github.com/alielmi98/tmdb-cli-tool/internal/models"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tmdb-app",
	Short: "A CLI tool to fetch data from The Movie Database (TMDB)",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Please provide a valid command.")
	},
}

func init() {
	rootCmd.AddCommand(playingCmd)
	rootCmd.AddCommand(popularCmd)
	rootCmd.AddCommand(topCmd)
	rootCmd.AddCommand(upcomingCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var playingCmd = &cobra.Command{
	Use:   "playing",
	Short: "Fetch now playing movies",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Fetching now playing movies...")
		moviesRes, err := api.FetchMovies("now_playing")
		if err != nil {
			fmt.Println(err)
			return
		}
		outputMovies(moviesRes)
	},
}

var popularCmd = &cobra.Command{
	Use:   "popular",
	Short: "Fetch popular movies",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Fetching popular movies...")
		moviesRes, err := api.FetchMovies("popular")
		if err != nil {
			fmt.Println(err)
			return
		}
		outputMovies(moviesRes)

	},
}

var topCmd = &cobra.Command{
	Use:   "top",
	Short: "Fetch top-rated movies",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Fetching top-rated movies...")
		moviesRes, err := api.FetchMovies("top_rated")
		if err != nil {
			fmt.Println(err)
			return
		}
		outputMovies(moviesRes)

	},
}

var upcomingCmd = &cobra.Command{
	Use:   "upcoming",
	Short: "Fetch upcoming movies",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Fetching upcoming movies...")
		moviesRes, err := api.FetchMovies("upcoming")
		if err != nil {
			fmt.Println(err)
			return
		}
		outputMovies(moviesRes)
	},
}

func outputMovies(moviesRes *models.MovieRes) {
	for key, value := range moviesRes.Results {
		fmt.Println("╔══════════════════════════════════════════════════╗")
		fmt.Printf("║  🎬  %d. %s\n", key+1, value.Title)
		fmt.Println("╠══════════════════════════════════════════════════╣")
		fmt.Printf("║  📅  Release Date: %s\n", value.ReleaseDate)
		fmt.Printf("║  ⭐  Vote Count: %d\n", value.VoteCount)
		fmt.Printf("║  ⭐  Rating: %.1f\n", value.VoteAverage)
		fmt.Println("╠══════════════════════════════════════════════════╣")
		fmt.Printf("║  📖  Overview: %s\n", value.Overview)
		fmt.Println("╚══════════════════════════════════════════════════╝")
		fmt.Println()
	}
}
