import SwiftUI

struct Note: Identifiable {
    let id = UUID()
    let text: String
    let date: String
    let color: Color
    let imageName: String
    let uploadedImage: UIImage? // New property for the uploaded image
}
