import SwiftUI
import PhotosUI


struct NoteView: View {
    @State private var selectedColor: Color = .blue.opacity(0.5)
    @State private var noteText: String = ""
    @State private var dateText: String = ""
    @State private var showDatePicker: Bool = false
    @State private var selectedDate = Date()
    @State private var navigateToNoteList = false
    @State private var uploadedImage: UIImage? = nil // State to hold the uploaded image
    @State private var showingImagePicker = false // To show the image picker
    @Environment(\.presentationMode) var presentationMode
    @EnvironmentObject var noteStore: NoteStore // Using environment object for shared data
    
    let colors: [(Color, String)] = [
        (.blue.opacity(0.5), "冷静"),
        (.orange, "担忧"),
        (.blue, "难过"),
        (.yellow, "愤怒"),
        (.red, "开心"),
        (.green.opacity(0.5), "害怕"),
        (.mint, "期待"),
        (.purple.opacity(0.7), "焦虑")
    ]
    
    var body: some View {
        NavigationView {
            VStack(spacing: 20) {
                HStack {
                    Button(action: {
                        presentationMode.wrappedValue.dismiss()
                    }) {
                        Image(systemName: "chevron.left")
                            .font(.title)
                            .padding()
                    }
                    Spacer()
                    Button(action: {
                        showingImagePicker = true // Show the image picker
                    }) {
                        Image(systemName: "camera")
                            .font(.title)
                            .padding()
                    }
                }
                
                VStack {
                    Text("记录心情瞬间")
                        .foregroundColor(.black)
                        .frame(maxWidth: .infinity, alignment: .leading)
                        .padding(.leading)
                    
                    if let image = uploadedImage {
                        Image(uiImage: image)
                            .resizable()
                            .scaledToFit()
                            .frame(height: 150)
                            .padding()
                            .background(Color.gray.opacity(0.2))
                            .cornerRadius(10)
                            .onLongPressGesture {
                                uploadedImage = nil // Remove the image on long press
                            }
                    }
                    
                    TextEditor(text: $noteText)
                        .foregroundColor(.black)
                        .frame(maxWidth: .infinity, minHeight: 100)
                        .padding()
                        .background(selectedColor)
                        .cornerRadius(10)
                        .padding(.horizontal)
                    
                    Text("选取心情颜色")
                        .foregroundColor(.black)
                        .frame(maxWidth: .infinity, alignment: .leading)
                        .padding(.leading)
                    
                    LazyVGrid(columns: Array(repeating: GridItem(.flexible()), count: 4)) {
                        ForEach(colors, id: \.1) { color, label in
                            VStack {
                                Circle()
                                    .fill(color)
                                    .frame(width: 50, height: 50)
                                    .onTapGesture {
                                        selectedColor = color
                                    }
                                Text(label)
                                    .font(.caption)
                                    .foregroundColor(.black)
                            }
                        }
                    }
                    .padding(.horizontal)
                    
                    TextField("记录一下日期", text: $dateText)
                        .textFieldStyle(RoundedBorderTextFieldStyle())
                        .padding(.horizontal)
                        .onTapGesture {
                            showDatePicker.toggle()
                        }
                    
                    if showDatePicker {
                        DatePicker("选择日期", selection: $selectedDate, displayedComponents: .date)
                            .datePickerStyle(GraphicalDatePickerStyle())
                            .onChange(of: selectedDate) { newDate in
                                dateText = formattedDate(newDate)
                            }
                            .padding(.horizontal)
                    }
                    
                    Button(action: {
                        let imageName = imageForColor(selectedColor)
                        let newNote = Note(
                            text: noteText,
                            date: dateText,
                            color: selectedColor,
                            imageName: imageName,
                            uploadedImage: uploadedImage // Save the uploaded image
                        )
                        noteStore.notes.append(newNote)
                        presentationMode.wrappedValue.dismiss()
                    }) {
                        Text("完成")
                            .foregroundColor(.white)
                            .padding()
                            .frame(maxWidth: .infinity)
                            .background(
                                LinearGradient(
                                    gradient: Gradient(colors: [
                                        Color(red: 255/255, green: 222/255, blue: 222/255).opacity(0.8),
                                        Color(red: 253/255, green: 173/255, blue: 171/255).opacity(0.8)
                                    ]),
                                    startPoint: .top,
                                    endPoint: .bottom
                                ))
                            .cornerRadius(10)
                            .padding(.horizontal)
                    }
                    
                    NavigationLink(destination: NoteListView().navigationBarBackButtonHidden(true), isActive: $navigateToNoteList) {
                        Button(action: {
                            navigateToNoteList = true
                        }) {
                            Text("我的心情储存罐")
                                .foregroundColor(.pink.opacity(0.5))
                                .padding()
                                .frame(maxWidth: .infinity)
                                .background(Color.white)
                                .overlay(
                                    RoundedRectangle(cornerRadius: 10)
                                        .stroke(Color.pink.opacity(0.5), lineWidth: 2)
                                )
                                .padding(.horizontal)
                        }
                    }
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
            .sheet(isPresented: $showingImagePicker) {
                ImagePicker(selectedImage: $uploadedImage) // Show image picker
            }
        }
    }
    
    func formattedDate(_ date: Date) -> String {
        let formatter = DateFormatter()
        formatter.dateStyle = .medium
        return formatter.string(from: date)
    }
    
    func imageForColor(_ color: Color) -> String {
        switch color {
        case .red: return "image1"
        case .purple.opacity(0.7): return "image2"
        case .green.opacity(0.5), .mint: return "image3"
        case .blue.opacity(0.5), .blue: return "image4"
        case .yellow, .orange: return "image5"
        default: return "defaultImage"
        }
    }
}
struct NoteView_Previews: PreviewProvider {
    static var previews: some View {
        NoteView()
    }
}
