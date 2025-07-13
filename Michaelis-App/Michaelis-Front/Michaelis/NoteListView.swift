import SwiftUI

struct NoteListView: View {
    @EnvironmentObject var noteStore: NoteStore // 使用环境对象共享数据
    @State private var selectedNote: Note? = nil
    @State private var showingNoteReadView = false
    @State private var showDeleteAlert = false
    @State private var noteToDelete: Note?
    @Environment(\.presentationMode) var presentationMode
    
    
    var body: some View {
        VStack {
            // Top Navigation Bar
            HStack {
                  Button(action: {
                    presentationMode.wrappedValue.dismiss()
                  }) {
                    Image(systemName: "chevron.left")
                    .font(.title)
                    .padding(.leading)
                 }
                Spacer()
                Text("心情存储罐")
                .font(.headline)
                Spacer()
                Button(action: {
                    // More options button action
                }) {
                Image(systemName: "ellipsis")
                .font(.title)
                .padding(.trailing)
                           }
                       }
                       .padding(.top, 50)
            //Notes Grid
            LazyVGrid(columns: Array(repeating: GridItem(.flexible(), spacing: 20), count: 3), spacing: 20) {
                ForEach(noteStore.notes) { note in
                    NavigationLink(destination: NoteReadView(note: note)) {
                        VStack {
                            Image(note.imageName)
                                .resizable()
                                .aspectRatio(contentMode: .fit)
                                .frame(width: 80, height: 80)
                                .onLongPressGesture {
                                    noteToDelete = note
                                    showDeleteAlert = true
                                }
                            Text(note.date)
                                .font(.caption)
                                .foregroundColor(.gray)
                        }
                    }
                }
            }

                        .padding()
                        .alert(isPresented: $showDeleteAlert) {
                            Alert(
                                title: Text("确认删除"),
                                message: Text("确定要删除这个心情记录吗？"),
                                primaryButton: .destructive(Text("删除")) {
                                    if let noteToDelete = noteToDelete {
                                        noteStore.delete(noteToDelete)
                                    }
                                },
                                secondaryButton: .cancel()
                            )
                        }
                        
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
        .edgesIgnoringSafeArea(.all)
        
    }
    
    private func selectNoteForReading(note: Note) {
        selectedNote = note
        showingNoteReadView = true
        print("Selected note: \(String(describing: selectedNote))")
    }
    
}

struct NoteListView_Previews: PreviewProvider {
    static var previews: some View {
        NoteListView().environmentObject(NoteStore())
    }
}
