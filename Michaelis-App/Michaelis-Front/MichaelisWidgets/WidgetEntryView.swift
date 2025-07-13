import SwiftUI
import WidgetKit

struct WidgetEntryView: View {
    var entry: SimpleEntry

    var body: some View {
        ZStack {
            // 设置背景颜色为线性渐变
            RoundedRectangle(cornerRadius: 15)
                .fill(LinearGradient(gradient: Gradient(colors: [
                    Color(red: 218/255, green: 250/255, blue: 215/255, opacity: 1),
                    Color(red: 245/255, green: 236/255, blue: 208/255, opacity: 1),
                    Color(red: 240/255, green: 247/255, blue: 220/255, opacity: 1),
                    Color(red: 242/255, green: 219/255, blue: 216/255, opacity: 1)
                ]), startPoint: .top, endPoint: .bottom))
                .edgesIgnoringSafeArea(.all)
            
            VStack(alignment: .leading, spacing: 8) {
                // 显示主要内容的文字
                Text("到目前为止，这一周过得如何？")
                    .font(.system(size: 12, weight: .bold))
                    .foregroundColor(.black)
                
                Text("今天完成了什么值得骄傲的事情？")
                    .font(.system(size: 12))
                    .foregroundColor(.black)
                
                Text("今天有什么值得感恩的事情吗？")
                    .font(.system(size: 12))
                    .foregroundColor(.black)
                
                Text("有什么值得记录的开心的事情？")
                    .font(.system(size: 12))
                    .foregroundColor(.black)
                
                Spacer()
                
                // 添加头像和文本部分
                HStack {
                    // 头像图片
                    Image("dulaAvatar") // 确保你已经将图像添加到项目的 Assets 中
                        .resizable()
                        .scaledToFit()
                        .frame(width: 50, height: 50)
                        .clipShape(Circle()) // 将图片裁剪为圆形
                    
                    // 显示在头像旁边的文字
                    Text("我一直都在，来记录今日份心情吧!")
                        .font(.system(size: 12))
                        .foregroundColor(.black)
                        .padding(8)
                        .background(RoundedRectangle(cornerRadius: 8).fill(Color.white.opacity(0.7)))
                }
            }
            .padding(16)
        }
    }
}


