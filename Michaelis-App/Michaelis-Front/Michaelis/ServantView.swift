import SwiftUI

struct ServantView: View {
    @State private var imageName = "wizard"
    @State private var chatText = "Hi! 小主人，希望您昨晚休息得好!"
    @Environment(\.presentationMode) var presentationMode
    var body: some View {
        VStack {
            HStack {
                Button(action: {
                    presentationMode.wrappedValue.dismiss()
                }) {
                    Image(systemName: "chevron.left")
                        .font(.title)
                        .padding(.leading)
                }
                Spacer()
                Text("我的管家")
                    .font(.headline)
                    .padding()
                Spacer()
                Button(action: {
                    // Share button action
                }) {
                    Image(systemName: "ellipsis")
                        .font(.title)
                        .padding(.trailing)
                }
            }
            .padding(.top, 10)
            Spacer()
            
            VStack(spacing: 10) {
                ChatBubble2(text: chatText, isReceived: true)
                ChatBubble2(text: "我会一直都在，请期待好事发生!", isReceived: false)
            }
            .padding()
            
            ZStack {
                Image(imageName)
                    .resizable()
                    .aspectRatio(contentMode: .fit)
                    .frame(height: 250)
                
                VStack {
                    Spacer()
                    HStack {
                        Button(action: {
                            imageName = "duola"
                            chatText = "嗨！大雄，让我用我的四维口袋里的道具来帮你解决问题吧！"
                        }) {
                            
                            Image("change")
                                .resizable()
                                .frame(width: 160, height: 40)
                        }
                        .padding(.leading, 20)
                        Spacer()
                    }
                    .padding(.bottom, 50)
                }
                .padding(-25.0)
            }
            
            Spacer()
            
            HStack(spacing: 50) {
                NavigationLink(destination: SettingView().navigationBarBackButtonHidden(true)) {
                    VStack {
                        Image("moreservice")
                            .resizable()
                            .frame(width: 80, height: 50)
                        Text("更多服务")
                            .foregroundColor(Color.black)
                    }
                }
                NavigationLink(destination: ChatView().navigationBarBackButtonHidden(true)) {
                                    VStack {
                                        Image("chat")
                                            .resizable()
                                            .frame(width: 80, height: 50)
                                        Text("聊天")
                                            .foregroundColor(Color.black)
                                    }
                                }
            }
            .padding()
        }
        .padding()
        .background(LinearGradient(gradient: Gradient(colors: [Color(red: 255/255, green: 255/255, blue: 232/255, opacity: 1), Color(red: 217/255, green: 255/255, blue: 224/255, opacity: 1)]), startPoint: .top, endPoint: .bottom))
    }
}

struct ChatBubble2: View {
    var text: String
    var isReceived: Bool
    
    var body: some View {
        HStack {
            if isReceived {
                Text(text)
                    .padding()
                    .background(LinearGradient(gradient: Gradient(colors: [Color(red: 255/255, green: 176/255, blue: 179/255, opacity: 0.8), Color(red: 250/255, green: 208/255, blue: 196/255, opacity: 0.8)]), startPoint: .leading, endPoint: .trailing))
                    .cornerRadius(10)
                    .shadow(color: Color.black.opacity(0.25), radius: 2, x: 4, y: 4)
                Spacer()
            } else {
                Spacer()
                Text(text)
                    .padding()
                    .background(LinearGradient(gradient: Gradient(colors: [Color(red: 255/255, green: 176/255, blue: 179/255, opacity: 0.8), Color(red: 250/255, green: 208/255, blue: 196/255, opacity: 0.8)]), startPoint: .leading, endPoint: .trailing))
                    .cornerRadius(10)
                    .shadow(color: Color.black.opacity(0.25), radius: 2, x: 4, y: 4)
            }
        }
    }
}

struct ServantView_Previews: PreviewProvider {
    static var previews: some View {
        ServantView()
    }
}
