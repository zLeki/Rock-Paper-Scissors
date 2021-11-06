package main 
import (
  "fmt"
  "github.com/bwmarrin/discordgo"
  "os"
	"os/signal"
	"syscall"
  "math/rand"
)
var originaluserid = ""
func main() {
     dg, err := discordgo.New("Bot " + "token")
    if err != nil {
        fmt.Println("error created while making a bot")
        return
    }
    dg.AddHandler(on_message)
    dg.AddHandler(on_reaction)
    err = dg.Open()
    if err != nil {
        fmt.Println("Error created while opening the bot", err)
        return
    }
    fmt.Println("Bot is up and running :sunglasses:")
    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
    <-sc
    
}
func on_message(s *discordgo.Session, m *discordgo.MessageCreate) {
    
    if m.Author.ID != s.State.User.ID {
        if m.Content == ".rps" {
            originaluserid = m.Author.ID
            msg, err := s.ChannelMessageSend(m.ChannelID, "Rock, Paper, Or Scissors?")
            if err != nil {
                fmt.Println("error when sending message", err)
                return
            }
            s.MessageReactionAdd(m.ChannelID, msg.ID, "🗿")
            s.MessageReactionAdd(m.ChannelID, msg.ID, "📄")
            s.MessageReactionAdd(m.ChannelID, msg.ID, "✂️")
        }
    }
}
func on_reaction(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
    if r.UserID != s.State.User.ID {
      fmt.Println(originaluserid)
      if r.UserID == originaluserid {
        
        reasons := make([]string, 0)
        var paper = ":page_facing_up:"
        reasons = append(reasons,
                "rock", // rock
                "paper", // paper
                "scissors") // scissors
        selected := reasons[rand.Intn(len(reasons))]
        
        if r.Emoji.Name == "✂️" {
            if selected != "scissors" {
                if selected == "rock" {
                    s.ChannelMessageSend(r.ChannelID, r.Emoji.Name+":left_facing_fist:"+":"+selected+":")
                    s.ChannelMessageSend(r.ChannelID, "You lose :skull:")
                }else if selected == "paper" {
                    s.ChannelMessageSend(r.ChannelID, r.Emoji.Name+":right_facing_fist:"+paper)
                    s.ChannelMessageSend(r.ChannelID, "You win :tada:")
                }
            }else {
                s.ChannelMessageSend(r.ChannelID, r.Emoji.Name+":handshake:"+":"+selected+":")
                s.ChannelMessageSend(r.ChannelID, `"Fair trade" - Drake`)
            }
            
        }else if r.Emoji.Name == "📄" {
            if selected != "paper" {
                if selected == "scissors" {
                    s.ChannelMessageSend(r.ChannelID, r.Emoji.Name+":left_facing_fist:"+":"+selected+":")
                    s.ChannelMessageSend(r.ChannelID, "You lose :skull:")
                }else if selected == "rock" {
                    s.ChannelMessageSend(r.ChannelID, r.Emoji.Name+":right_facing_fist:"+":"+selected+":")
                    s.ChannelMessageSend(r.ChannelID, "You win :tada:")
                }
            }else {
                s.ChannelMessageSend(r.ChannelID, r.Emoji.Name+":handshake:"+paper)
                s.ChannelMessageSend(r.ChannelID, `"Fair trade" - Drake`)
            }
        
        }else if r.Emoji.Name == "🗿" {
            if selected != "rock" {
                if selected == "paper" {
                    s.ChannelMessageSend(r.ChannelID, r.Emoji.Name+":left_facing_fist:"+paper)
                    s.ChannelMessageSend(r.ChannelID, "You lose :skull:")
                }else if selected == "scissors" {
                    s.ChannelMessageSend(r.ChannelID, r.Emoji.Name+":right_facing_fist:"+":"+selected+":")
                    s.ChannelMessageSend(r.ChannelID, "You win :tada:")
                }
            }else {
                s.ChannelMessageSend(r.ChannelID, r.Emoji.Name+":handshake:"+":moyai:")
                s.ChannelMessageSend(r.ChannelID, `"Fair trade" - Drake`)
            }
          originaluserid = ""
        }
    }
    }
}
