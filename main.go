package main
import(
	"fmt"
	"syscall"
	"os"
	"os/signal"
	"github.com/bwmarrin/discordgo"
)
const(
	prefix = "go!"
	token = "ODcwNjA2MzM1Mjg3Mzg2MTIy.YQPNVQ.GWfF2yehUN8oKUtWSVEiM7DlA7s"
)
func main() {
	fmt.Println("Proje Başarıyla başlatıldı.")
	client, err := discordgo.New("Bot " + token);
	if err != nil {
		fmt.Println(err)
		return
	}
	client.AddHandler(messageCreate)
	client.Identify.Intents = discordgo.IntentsGuildMessages
	err = client.Open()
	if err != nil { 
		fmt.Println("Başlatılırken hata oluştu: ", err)
		return
	}
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	client.Close()
}
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "Sa" {
		s.ChannelMessageSend(m.ChannelID, "<@" + m.Author.ID + ">, Aleykümselam. Hoşgeldin.")
	}
	if m.Content == prefix + "react" {
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title: "Tepki Ekleme İşlemi",
			Description: "**Başarılı**\nBaşarıyla Mesajınıza 📀 tepkisi eklendi!",
			Color: 12430073,
		})
		s.MessageReactionAdd(m.ChannelID, m.Reference().MessageID, "📀")
	}
}