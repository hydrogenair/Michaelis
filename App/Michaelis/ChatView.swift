import SwiftUI
import GoogleGenerativeAI

struct ChatView: View {
    let model = GenerativeModel(name: "gemini-pro", apiKey: APIKey.default)
    @State private var message: String = ""
    @State private var messages: [Message] = [
        Message(text: "大雄，你已经坐了好久了哦！要不要站起来活动一下？动动身体，伸展伸展！", isReceived: true),
        Message(text: "大雄，你回来啦!感觉你心情不好，怎么啦？", isReceived: true)
    ]
    @State private var showTypingIndicator: Bool = false
    @Environment(\.presentationMode) var presentationMode
    
    var body: some View {
        GeometryReader { geometry in
            ZStack {
                Image("duola")
                    .resizable()
                    .aspectRatio(contentMode: .fill)
                    .frame(width: geometry.size.width)
                
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
                        Image("dulaAvatar")
                            .resizable()
                            .frame(width: 50, height: 50)
                            .clipShape(Circle())
                            .padding(.trailing)
                        Spacer()
                    }
                    .padding(.top, 10)
                    
                    Divider()
                    
                    VStack {
                        ScrollView {
                            VStack(spacing: 10) {
                                ForEach(messages) { message in
                                    if message.isReceived {
                                        ReceivedMessage(text: message.text)
                                    } else {
                                        SentMessage(text: message.text)
                                    }
                                }
                                if showTypingIndicator {
                                    TypingIndicator()
                                }
                            }
                            .padding(25.0)
                        }
                        
                        Divider()
                        
                        HStack {
                            TextField("您可以向我倾诉任何事……", text: $message)
                                .padding()
                                .background(Color.white)
                                .cornerRadius(10)
                            
                            Button(action: {
                                sendMessage()
                            }) {
                                Text("send")
                                    .foregroundColor(Color.white)
                                    .padding()
                                    .background(Color.pink.opacity(0.2))
                                    .clipShape(Circle())
                            }
                        }
                        .padding()
                        .background(LinearGradient(gradient: Gradient(colors: [Color(red: 255/255, green: 176/255, blue: 179/255, opacity: 0.8), Color(red: 250/255, green: 208/255, blue: 196/255, opacity: 0.8)]), startPoint: .leading, endPoint: .trailing))
                        .cornerRadius(20)
                        .shadow(radius: 5)
                        .padding()
                    }
                    .background(
                        Color.white.opacity(0.7)
                            .cornerRadius(20)
                            .shadow(radius: 10)
                            .padding()
                    )
                }
            }
        }
    }
    
    func sendMessage() {
        guard !message.isEmpty else { return }
        
        // Add user's message to the chat
        let sentMessage = Message(text: message, isReceived: false)
        messages.append(sentMessage)
        let userMessage = message
        message = ""
        
        // Show typing indicator
        showTypingIndicator = true
        
        // Send the user's message to Google Generative AI
        Task {
            do {
                let prompt = "模仿多啦a梦的语气与我对话，带上多啦a梦特有的惯用词，每次只回答一句话，根据我说的话给予我心里抚慰，对我的称呼是大雄，现在我说“\(userMessage)”"
                let response = try await model.generateContent(prompt)
                
                if let text = response.text {
                    // Remove typing indicator
                    showTypingIndicator = false
                    
                    // Append AI response to chat
                    let receivedMessage = Message(text: text, isReceived: true)
                    messages.append(receivedMessage)
                }
            } catch {
                print("Error generating response: \(error)")
                showTypingIndicator = false
            }
        }
    }
}

struct ReceivedMessage: View {
    var text: String
    
    var body: some View {
        HStack(alignment: .top) {
            Image("dulaAvatar")
                .resizable()
                .frame(width: 40, height: 40)
                .clipShape(Circle())
            Text(text)
                .padding()
                .background(Color(red: 255/255, green: 255/255, blue: 237/255, opacity: 1))
                .cornerRadius(10)
                .shadow(color: Color.black.opacity(0.25), radius: 2, x: 4, y: 4)
            Spacer()
        }
    }
}

struct SentMessage: View {
    var text: String
    
    var body: some View {
        HStack(alignment: .top) {
            Spacer()
            Text(text)
                .foregroundColor(Color.white)
                .padding()
                .background(Color(red: 252/255, green: 136/255, blue: 140/255, opacity: 1))
                .cornerRadius(10)
                .shadow(color: Color.black.opacity(0.25), radius: 2, x: 4, y: 4)
            Image("avatar")
                .resizable()
                .frame(width: 40, height: 40)
                .clipShape(Circle())
        }
    }
}

struct TypingIndicator: View {
    @State private var dotCount = 1
    
    var body: some View {
        HStack(alignment: .top) {
            Image("dulaAvatar")
                .resizable()
                .frame(width: 40, height: 40)
                .clipShape(Circle())
            Text("...".prefix(dotCount))
                .padding()
                .background(Color(red: 255/255, green: 255/255, blue: 237/255, opacity: 1))
                .cornerRadius(10)
                .shadow(color: Color.black.opacity(0.25), radius: 2, x: 4, y: 4)
            Spacer()
        }
        .onAppear {
            Timer.scheduledTimer(withTimeInterval: 0.5, repeats: true) { timer in
                dotCount = (dotCount % 3) + 1
            }
        }
    }
}

struct Message: Identifiable {
    let id = UUID()
    var text: String
    var isReceived: Bool
}

struct ChatView_Previews: PreviewProvider {
    static var previews: some View {
        ChatView()
    }
}
