import SwiftUI

struct FlowerView: View {
    var flowerImage: String
    var message: String
    @Environment(\.presentationMode) var presentationMode
    var body: some View {
        ZStack {
            // Background Image
            Image("emoser_back")
                .resizable()
                .scaledToFill()
                .edgesIgnoringSafeArea(.all)
            
            VStack {
                // Navigation Bar with Back Button and Title
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
                }
                .padding(.horizontal, 20)
                .padding(.top, 10)
                
                Spacer()
                
                // Doraemon Character
                Image("dla")
                    .resizable()
                    .frame(width: 100, height: 120)
                    .padding(.top, -30) // Adjust position if necessary
                
                // Message Bu1bble with Text
                Text(message)
                    .padding()
                    .background(Color.pink.opacity(0.6))
                    .cornerRadius(15)
                    .foregroundColor(.white)
                    .font(.body)
                    .multilineTextAlignment(.center)
                    .padding(.horizontal, 20)
                    .padding(.top, 10)
                
                // Flower Image
                Image(flowerImage) // Replace with the correct image name
                    .resizable()
                    .scaledToFit()
                    .frame(width: 160, height: 120)
                    .background(Color.white.opacity(0.5)) // White overlay for effect
                    .cornerRadius(15)
                    .padding(.top, 10)
                
                Spacer()
            }
        }
    }
}

struct FlowerView_Previews: PreviewProvider {
    static var previews: some View {
        FlowerView(flowerImage: "flower5", message: "大雄，希望你的心情能够像太阳花一样灿烂起来～")
    }
}
