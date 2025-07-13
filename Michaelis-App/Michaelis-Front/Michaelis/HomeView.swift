import SwiftUI

struct HomeView: View {
    // Define the array of image and text pairs
    let imageTextPairs = [
        ("character", "Hi! 久坐不太好哦，尝试去放松一下吧!"),
        ("badminton", "好久没有和朋友约羽毛球了，今天试试吧"),
        ("bike", "今天阳光明媚，要不要外出骑行呢？"),
        ("book_read", "换换心情，久违的来读一本书吧"),
        ("life_records", "生活的精彩也可以随时记录哦"),
        ("pet", "宠物也可以治愈心情不好的一天"),
        ("phone_play", "不要长时间低头玩手机嗷"),
        ("sleep", "主人不要熬夜哦，今天休息时间小憩一会儿吧"),
        ("tennis", "好久没有和朋友约网球了，今天试试吧"),
        ("tired 1", "最近感觉身体疲惫，工作上有什么烦恼吗"),
        ("travel", "如果感到烦闷，不如试试说走就走的旅行"),
        ("upset_sweet", "让甜食扫荡坏心情～")
    ]

    // Randomly select an image-text pair
    @State private var selectedPair: (String, String) = ("character", "Hi! 久坐不太好哦，尝试去放松一下吧!")
    
    var body: some View {
        NavigationView {
            VStack(spacing: 60) {
                HStack(spacing: 30) {
                    NavigationLink(destination: ActivityView().navigationBarBackButtonHidden(true)) {
                        VStack {
                            Image("activity")
                                .resizable()
                                .frame(width: 45, height: 50)
                            Text("Activity")
                                .foregroundColor(Color.black)
                        }
                    }
                    NavigationLink(destination: HealthRecordsView().navigationBarBackButtonHidden(true)) {
                        VStack {
                            Image("records")
                                .resizable()
                                .frame(width: 50, height: 50)
                            Text("Health Records")
                                .foregroundColor(Color.black)
                        }
                    }
                    NavigationLink(destination: EmotionView().navigationBarBackButtonHidden(true)) {
                        VStack {
                            Image("emotion")
                                .resizable()
                                .frame(width: 50, height: 50)
                            Text("Emotion")
                                .foregroundColor(Color.black)
                        }
                    }
                }
                .font(.headline)
                .padding(.top, 80)
                
                ZStack(alignment: .top) {
                    HStack {
                        Spacer()
                        NavigationLink(destination: ServantView().navigationBarBackButtonHidden(true)) {
                            VStack(spacing: 10) {
                                Image("chatroom")
                                    .resizable()
                                    .frame(width: 50, height: 50)
                                Text("ChatRoom")
                                    .foregroundColor(Color.black)
                            }
                            .padding(.trailing, 20)
                        }
                    }
                    // Display selected random image
                    Image(selectedPair.0)
                        .resizable()
                        .scaledToFit()
                        .frame(height: 400)
                        .frame(maxWidth: .infinity, alignment: .top)
                }

                // Display selected random text
                ChatBubble(text: selectedPair.1)
                    .frame(width: 345, height: 63)
                    .padding(.bottom, 10)

                Spacer()
            }
            .onAppear {
                // Select a new random image-text pair when the view appears
                selectedPair = imageTextPairs.randomElement() ?? ("character", "Hi! 久坐不太好哦，尝试去放松一下吧!")
            }
            .navigationBarHidden(true)
            .background(
                LinearGradient(
                    gradient: Gradient(colors: [
                        Color(red: 255 / 255, green: 255 / 255, blue: 232 / 255, opacity: 1),
                        Color(red: 217 / 255, green: 255 / 255, blue: 224 / 255, opacity: 1)
                    ]),
                    startPoint: .top,
                    endPoint: .bottom
                )
            )
        }
    }
}

struct ChatBubble: View {
    var text: String
    
    var body: some View {
        ZStack {
            RoundedRectangle(cornerRadius: 15)
                .fill(
                    LinearGradient(
                        gradient: Gradient(colors: [
                            Color(red: 1, green: 0.69, blue: 0.7),
                            Color(red: 0.98, green: 0.81, blue: 0.77),
                            Color(red: 0.98, green: 0.81, blue: 0.77)
                        ]),
                        startPoint: UnitPoint(x: 1, y: 0.5),
                        endPoint: UnitPoint(x: 0.45773, y: 0.5)
                    )
                )
                .shadow(color: Color.black.opacity(0.25), radius: 4, x: 4, y: 4)
            
            Text(text)
                .padding()
                .background(Color.clear)
                .cornerRadius(15)
        }
        .frame(width: 345, height: 63)
    }
}

struct HomeView_Previews: PreviewProvider {
    static var previews: some View {
        HomeView()
    }
}
