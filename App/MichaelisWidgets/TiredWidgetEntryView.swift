import SwiftUI
import WidgetKit

// 新的 Widget 视图
struct TiredWidgetEntryView: View {
    var entry: SimpleEntry

    var body: some View {
        ZStack {
            // 背景颜色
            RoundedRectangle(cornerRadius: 15)
                .fill(LinearGradient(gradient: Gradient(colors: [
                    Color(red: 218/255, green: 250/255, blue: 215/255, opacity: 1),
                    Color(red: 245/255, green: 236/255, blue: 208/255, opacity: 1),
                    Color(red: 240/255, green: 247/255, blue: 220/255, opacity: 1),
                    Color(red: 242/255, green: 219/255, blue: 216/255, opacity: 1)
                ]), startPoint: .top, endPoint: .bottom))
                .edgesIgnoringSafeArea(.all)
            
            VStack {
                // 图片
                Image("tired") // 确保图片名为 tiredImage
                    .resizable()
                    .scaledToFit()
                    .frame(width: 100, height: 100)

                Spacer()
                
                // 文本
                VStack {
                    Text("久坐记得运动")
                        .font(.system(size: 8))
                        .foregroundColor(.black)
                        .padding()
                        .background(RoundedRectangle(cornerRadius: 8).fill(Color.white.opacity(0.7)))
                }
            }
            .padding()
        }
    }
}
