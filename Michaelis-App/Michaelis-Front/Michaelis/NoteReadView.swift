import SwiftUI

struct NoteReadView: View {
    var note: Note
    @Environment(\.presentationMode) var presentationMode
    
    var body: some View {
        VStack(spacing: 20) {
            // 标题
            Text("记录内容")
                .font(.title2)
                .fontWeight(.bold)
                .padding(.top, 20)
            
            // 如果有上传的图片，显示图片；如果没有，则显示默认logo
            if let uploadedImage = note.uploadedImage {
                Image(uiImage: uploadedImage)
                    .resizable()
                    .scaledToFit()
                    .frame(maxWidth: 200, maxHeight: 200) // 设置图片尺寸
                    .cornerRadius(10)
                    .padding(.horizontal, 20)
            } else {
            }
            
            // 内容框
            ScrollView {
                Text(note.text)
                    .padding()
                    .frame(maxWidth: .infinity)
                    .frame(height: 500) // 固定高度
                    .background(RoundedRectangle(cornerRadius: 10).fill(note.color))
                    .padding(.horizontal, 20)
            }
            .frame(maxHeight: 500) // 限制滚动框的最大高度

            // 显示记录日期
            Text("记录日期: \(note.date)")
                .font(.caption)
                .foregroundColor(.gray)
            
            Spacer()
        }
        .frame(maxWidth: .infinity, maxHeight: .infinity)
        .background(
            LinearGradient(
                gradient: Gradient(colors: [
                    Color(red: 255/255, green: 255/255, blue: 232/255),
                    Color(red: 217/255, green: 255/255, blue: 224/255)
                ]),
                startPoint: .top,
                endPoint: .bottom
            )
            .ignoresSafeArea()
        )
        .navigationBarBackButtonHidden(true)
        .navigationBarItems(leading: Button(action: {
            presentationMode.wrappedValue.dismiss()
        }) {
            Image(systemName: "chevron.left")
                .font(.title)
        })
    }
}

struct NoteReadView_Previews: PreviewProvider {
    static var previews: some View {
        let sampleNote = Note(text: "这是一个示例内容", date: "Nov 5, 2024", color: .blue, imageName: "image1", uploadedImage: nil) // 没有上传图片
        NoteReadView(note: sampleNote)
    }
}
