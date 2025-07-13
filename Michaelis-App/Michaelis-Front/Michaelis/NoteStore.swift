import SwiftUI
class NoteStore: ObservableObject {
    @Published var notes: [Note]
    init() {
           // 预先加载一些示例数据
           self.notes = [
            Note(text: "新学期转到其他学院，好多陌生人", date: "Nov 1, 2024", color:.green.opacity(0.5),imageName: "image3",uploadedImage: nil),
               Note(text: "上课迟到了，被批评了", date: "Nov 2, 2024", color:.blue,imageName: "image4",uploadedImage: nil),
               Note(text: "下周去游乐园～", date: "Nov 3, 2024",color:.mint, imageName: "image3",uploadedImage: nil)
           ]
       }
       

    func delete(_ note: Note) {
        notes.removeAll { $0.id == note.id }
    }
}
