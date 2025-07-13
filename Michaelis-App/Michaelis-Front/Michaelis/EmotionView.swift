import SwiftUI
import GoogleGenerativeAI

struct EmotionView: View {
    @State private var selectedDate = Date()
    @Environment(\.presentationMode) var presentationMode
    @EnvironmentObject var noteStore: NoteStore
    let model = GenerativeModel(name: "gemini-1.5-flash", apiKey: APIKey.default)
    
    var body: some View {
        VStack {
            // 顶部信息部分
            HStack {
                Button(action: {
                    presentationMode.wrappedValue.dismiss()
                }) {
                    Image(systemName: "chevron.left")
                        .font(.title)
                }
                Spacer()
                Button(action: {
                    // 更多选项按钮操作
                }) {
                    Image(systemName: "ellipsis")
                        .font(.title)
                }
            }
            .padding(.top, 50) // 向下移动图表以避免遮挡
            
            // 标题部分
            VStack(alignment: .leading, spacing: 10) {
                Text("专属管家提醒：")
                    .font(.headline)
                    .padding(.top)
                Text("主人，您昨天好像心情不太好。我可以为您准备一杯您喜欢的茶，或者安排一些您喜欢的活动，来缓解一下心情。")
                    .font(.subheadline)
                    .foregroundColor(.secondary)
            }
            .padding()
            .background(
                LinearGradient(gradient: Gradient(colors: [
                    Color(red: 255/255, green: 176/255, blue: 179/255, opacity: 1),
                    Color(red: 255/255, green: 226/255, blue: 222/255, opacity: 1),
                    Color(red: 250/255, green: 208/255, blue: 196/255, opacity: 1),
                    Color(red: 255/255, green: 237/255, blue: 232/255, opacity: 1)
                ]), startPoint: .topLeading, endPoint: .bottomTrailing)
            )
            .cornerRadius(30)
            .padding(.horizontal)
            
            // “向我倾诉”组件
            NavigationLink(destination: ServantView().navigationBarBackButtonHidden(true)) {
                HStack {
                    Spacer()
                    Text("向我倾诉")
                        .font(.caption)
                        .foregroundColor(Color.white)
                        .padding(.horizontal)
                        .padding(.vertical, 4)
                        .background(Color.black)
                        .cornerRadius(10)
                    Image("chatroom")
                        .resizable()
                        .frame(width: 40, height: 40)
                }
                .padding(.horizontal)
            }
            
            // 按钮部分
            HStack {
                NavigationLink(destination: NoteView().navigationBarBackButtonHidden(true)) {
                    HStack {
                        Image("record").resizable().frame(width: 40, height: 30)
                        Text("记录心情")
                            .font(.caption)
                            .foregroundColor(Color.black)
                    }
                    .padding()
                    .background(
                        LinearGradient(
                            gradient: Gradient(colors: [
                                Color(red: 255/255, green: 222/255, blue: 222/255).opacity(0.25),
                                Color(red: 253/255, green: 173/255, blue: 171/255).opacity(0.25)
                            ]),
                            startPoint: .top,
                            endPoint: .bottom
                        )
                    )
                    .cornerRadius(12)
                }
                
                NavigationLink(destination: NoteListView().navigationBarBackButtonHidden(true)) {
                    HStack {
                        Image("delete").resizable().frame(width: 25, height: 30)
                        Text("情绪储存罐")
                            .font(.caption)
                            .foregroundColor(Color.black).padding(1.0)
                    }
                }
                .padding()
                .background(
                    LinearGradient(
                        gradient: Gradient(colors: [
                            Color(red: 255/255, green: 222/255, blue: 222/255).opacity(0.25),
                            Color(red: 253/255, green: 173/255, blue: 171/255).opacity(0.25)
                        ]),
                        startPoint: .top,
                        endPoint: .bottom
                    )
                )
                .cornerRadius(12)
            }
            .padding(.horizontal)
            
            // 情绪统计图表部分
            VStack(alignment: .leading) {
                Image("chart")
                    .resizable()
                    .frame(height: 300)
                HStack {
                    Text("情绪统计")
                        .font(.headline)
                    Spacer()
                    DatePicker("日期", selection: $selectedDate, displayedComponents: .date)
                        .labelsHidden()
                        .padding(6)
                        .background(Color.white.opacity(0.4))
                        .cornerRadius(12)
                }
                .padding(.horizontal)
            }
            .padding(.horizontal)
            
            // 情绪变化因素部分
            VStack {
                Text("-使我产生情绪变化的因素-")
                    .font(.caption)
                    .padding(.top)
                HStack {
                    MoodFactorView(imageview: "happy", color: .red)
                    MoodFactorView(imageview: "anger", color: .orange)
                    MoodFactorView(imageview: "anxiety", color: .purple.opacity(0.7))
                }
            }
            .padding(.horizontal)
            
            Spacer()
        }
        .background(
            LinearGradient(
                gradient: Gradient(colors: [
                    Color(red: 255/255, green: 255/255, blue: 232/255),
                    Color(red: 217/255, green: 255/255, blue: 224/255)
                ]),
                startPoint: .top,
                endPoint: .bottom
            )
        )
        .edgesIgnoringSafeArea(.top)
    }
}

struct MoodFactorView: View {
    var imageview: String
    var color: Color
    @EnvironmentObject var noteStore: NoteStore
    let model = GenerativeModel(name: "gemini-1.5-flash", apiKey: APIKey.default)
    
    @State private var keyword = "什么令我…？"
    @State private var showFullKeyword = false  // 控制是否展示完整的关键词
    
    var body: some View {
        VStack {
            Button(action: {
                showFullKeyword.toggle()  // 点击按钮时切换详细视图显示状态
            }) {
                Text(keyword)
                    .font(.caption)
                    .fontWeight(.regular)
                    .foregroundColor(color)
                    .padding()
                    .frame(width: 110) // 固定宽度
                    .background(Color.white)
                    .cornerRadius(10)
                    .shadow(color: Color.black.opacity(0.25), radius: 2, x: 2, y: 2)
            }
            .sheet(isPresented: $showFullKeyword) {  // 弹出详细视图
                VStack {
                    Text("完整关键词")
                        .font(.headline)
                        .padding()
                    Text(keyword)  // 显示完整关键词
                        .font(.subheadline)
                        .padding()
                }
                .presentationDetents([.medium])
            }
            
            Image(imageview)
                .resizable()
                .frame(width: 80, height: 80)
        }
        .padding(4)
        .background(color.opacity(0.1))
        .cornerRadius(10)
        .onAppear {
            Task {
                await generateKeyword(for: imageview)
            }
        }
    }
    
    func generateKeyword(for imageview: String) async {
        // 获取对应颜色的笔记内容和图片
        let noteData: (text: String, image: UIImage?)? = {
            switch imageview {
            case "anxiety":
                let note = noteStore.notes.first { $0.color == .purple.opacity(0.7) }
                return (note?.text ?? "", note?.uploadedImage)
            case "anger":
                let note = noteStore.notes.first { $0.color == .yellow }
                return (note?.text ?? "", note?.uploadedImage)
            case "happy":
                let note = noteStore.notes.first { $0.color == .red }
                return (note?.text ?? "", note?.uploadedImage)
            default:
                return nil
            }
        }()
        
        guard let noteText = noteData?.text, !noteText.isEmpty else { return }
        
        // 根据 imageview 设置不同的提示语
        let prompt: String = {
            switch imageview {
            case "anxiety":
                return "根据以下文本和图片用一个词总结出让我感到焦虑的事件：\(noteText)"
            case "anger":
                return "根据以下文本和图片用一个词总结出让我感到愤怒的事件：\(noteText)"
            case "happy":
                return "根据以下文本和图片用一个词总结出让我感到开心的事件：\(noteText)"
            default:
                return ""
            }
        }()
        
        do {
            // 根据是否有图片决定生成内容的方式
            if let image = noteData?.image {
                // 同时使用文本和图片生成
               let response = try await model.generateContent(prompt, image)
                if let generatedText = response.text {
                    keyword = generatedText
                }
            } else {
                // 仅使用文本生成
                let response = try await model.generateContent(prompt)
                // 获取生成的文本内容
                if let generatedText = response.text {
                    keyword = generatedText
                }
            }
            
           
            
        } catch {
            print("AI生成关键词出错：\(error)")
        }
    }

}

struct EmotionView_Previews: PreviewProvider {
    static var previews: some View {
        EmotionView().environmentObject(NoteStore())
    }
}
