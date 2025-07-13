import SwiftUI

struct EmotionWindowView: View {
    @State private var selectedFlower: String?
       @State private var selectedMessage: String?
    @Environment(\.presentationMode) var presentationMode
    @State private var isNavigationActive = false

       // Flower data
       let flowers = [
           ("flower1", "大雄，愿这朵玫瑰的热烈与浪漫，带给温暖的力量，让你的心情如花般绽放"),
           ("flower2", "大雄，愿你的生活如牡丹般绚烂多彩"),
           ("flower3", "大雄，愿你的心灵如茉莉般纯洁无瑕，让烦恼随风而去"),
           ("flower4", "大雄，愿你的生活如百合般纯洁和谐，每一天都充满希望与美好"),
           ("flower5", "大雄，希望你的心情能够像太阳花一样灿烂起来～"),
           ("flower6", "大雄，愿你的心情如郁金香般多彩多姿，每一天都充满活力与快乐")
       ]
    var body: some View {
        ZStack {
            // Background Image
            Image("emoser_back")
                .resizable()
                .scaledToFill()
                .edgesIgnoringSafeArea(.all)
            
            VStack {
                // Navigation Bar with Back Button, Title, and Edit Button
                HStack {
                    Button(action: {
                        presentationMode.wrappedValue.dismiss()
                    }) {
                        Image(systemName: "chevron.left")
                            .font(.title2)
                            .foregroundColor(.white)
                    }
                    Spacer()
                    Text("心情橱窗")
                        .font(.headline)
                        .foregroundColor(.white)
                    Spacer()
                    Button(action: {
                        // Edit action
                    }) {
                        Image(systemName: "pencil")
                            .font(.title2)
                            .foregroundColor(.white)
                    }
                }
                .padding(.horizontal, 20)
                .padding(.top, 10)
                
                Spacer()
                
                // Doraemon Message with Avatar
                HStack(alignment: .bottom) {
                    // Message Bubble
                    Text("主人，您的情绪评估为：低落\n送您一束花希望您的心情会变好~")
                        .padding()
                        .background(Color.pink.opacity(0.7))
                        .cornerRadius(15)
                        .foregroundColor(.white)
                        .font(.body)
                        .frame(maxWidth: .infinity, alignment: .leading)
                        .padding(.horizontal, 20)
                        .padding(.vertical, 10)
                    
                    // Doraemon Avatar
                    Image("dla")
                        .resizable()
                        .frame(width: 80, height: 100)
                        .offset(x: -10, y: 20)
                }
                .padding(.bottom, 15)
                
                // Flower Selection Title with Selection Frame
                VStack{
                    Text("选择一束自己喜欢的花吧")
                        .font(.headline)
                        .foregroundColor(.white)
                        .padding(.bottom, 10)
                        .padding(.top, 20)
                    
                    // Flower Selection Grid (2x3 layout)
                    LazyVGrid(columns: [GridItem(.flexible()), GridItem(.flexible())], spacing: 20) {
                        ForEach(0..<flowers.count, id: \.self) { index in
                            Button(action: {
                                                            selectedFlower = flowers[index].0
                                                            selectedMessage = flowers[index].1
                                isNavigationActive = true
                                                        }) {
                                                            Image(flowers[index].0)
                                                                .resizable()
                                                                .aspectRatio(contentMode: .fill)
                                                                .frame(width: 150, height: 120)
                                                                .background(Color.white.opacity(0.5))
                                                                .cornerRadius(10)
                                                        }
                                
                        }
                    
                }.padding(.horizontal, 20)
                 .padding(.bottom, 20)
                   
                }.background(Color.white.opacity(0.5))
                    .cornerRadius(10)
                Button(action: {
                    let randomFlower = flowers.randomElement()!
                                            selectedFlower = randomFlower.0
                                            selectedMessage = randomFlower.1
                                    
                    isNavigationActive = true
                }) {
                    Text("帮我选择")
                        .foregroundColor(.white)
                        .padding()
                        .frame(maxWidth: .infinity)
                        .background(Color.pink.opacity(0.7))
                        .cornerRadius(10)
                }.padding(.bottom, 10)
                
          
            }
            NavigationLink(
                destination: FlowerView(flowerImage: selectedFlower ?? "flower5", message: selectedMessage ?? "11").navigationBarBackButtonHidden(true),
                isActive: $isNavigationActive
            ) {
                EmptyView()
            }


        }
    }
}

struct EmotionWindowView_Previews: PreviewProvider {
    static var previews: some View {
        EmotionWindowView()
    }
}
