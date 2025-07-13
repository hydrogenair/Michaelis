import SwiftUI

struct TimerView: View {
    @Environment(\.presentationMode) var presentationMode
    @Binding var sportHours: Double
    @State private var timer: Timer?
    @State private var timeElapsed: TimeInterval = 0
    @State private var isRunning: Bool = true
    
    var body: some View {
        VStack {
            // Top bar with back and share buttons
            HStack {
                Button(action: {
                    presentationMode.wrappedValue.dismiss()
                }) {
                    Image(systemName: "chevron.left")
                        .font(.title)
                        .padding(.leading)
                }
                Spacer()
                Button(action: {
                    // Share button action
                }) {
                    Image(systemName: "square.and.arrow.up")
                        .font(.title)
                        .padding(.trailing)
                }
            }
            .padding(.top, 10)
            
            Spacer()
            HStack{
                Image(systemName: "figure.run").resizable()
                    .frame(width: 30, height: 40)
                Text("跑步")
            }
            
            // Timer display
            Text(timeString(from: timeElapsed))
                .font(.system(size: 80, weight: .bold, design: .monospaced))
                .padding()
                .background(
                    RoundedRectangle(cornerRadius: 20)
                        .fill(
                            LinearGradient(
                                gradient: Gradient(colors: [Color.blue.opacity(0.3), Color.purple.opacity(0.3)]),
                                startPoint: .topLeading,
                                endPoint: .bottomTrailing
                            )
                        )
                        .shadow(color: Color.black.opacity(0.2), radius: 10, x: 0, y: 10)
                )
                .padding()
            
            Spacer()
            
            // Control buttons
            HStack(spacing: 50) {
                Button(action: {
                    stopTimer()
                    // 将时间转换为小时并更新
                    let minutes = timeElapsed / 60
                    sportHours += minutes
                    presentationMode.wrappedValue.dismiss()
                }) {
                    Text("完成")
                        .font(.title)
                        .padding()
                        .background(Color(red: 255 / 255, green: 255 / 255, blue: 232 / 255, opacity: 1))
                        .foregroundColor(.black)
                        .cornerRadius(10)
                        .shadow(radius: 10)
                }
                
                Button(action: {
                    if isRunning {
                        pauseTimer()
                    } else {
                        resumeTimer()
                    }
                }) {
                    Text(isRunning ? "暂停" : "继续")
                        .font(.title)
                        .padding()
                        .background(
                            LinearGradient(gradient: Gradient(colors: [Color(red: 255/255, green: 102/255, blue: 102/255), Color(red: 255/255, green: 178/255, blue: 102/255)]), startPoint: .leading, endPoint: .trailing)
                        )
                        .foregroundColor(.white)
                        .cornerRadius(10)
                        .shadow(radius: 10)
                }
            }
            .padding()
        }
        .onAppear {
            startTimer()
        }
        .padding()
        .background(
            LinearGradient(gradient: Gradient(colors: [
                Color(red: 218/255, green: 250/255, blue: 215/255, opacity: 1),
                Color(red: 245/255, green: 236/255, blue: 208/255, opacity: 1),
                Color(red: 240/255, green: 247/255, blue: 220/255, opacity: 1),
                Color(red: 242/255, green: 219/255, blue: 216/255, opacity: 1)
            ]), startPoint: .top, endPoint: .bottom)
        )
    }
    
    func timeString(from time: TimeInterval) -> String {
        let minutes = Int(time) / 60
        let seconds = Int(time) % 60
        return String(format: "%02d:%02d", minutes, seconds)
    }
    
    func startTimer() {
        timer = Timer.scheduledTimer(withTimeInterval: 1, repeats: true) { _ in
            timeElapsed += 1
        }
    }
    
    func stopTimer() {
        timer?.invalidate()
        timer = nil
        isRunning = false
    }
    
    func pauseTimer() {
        timer?.invalidate()
        isRunning = false
    }
    
    func resumeTimer() {
        startTimer()
        isRunning = true
    }
}


struct TimerView_Previews: PreviewProvider {
    @State static var sportHours: Double = 0.0
    
    static var previews: some View {
        TimerView(sportHours: $sportHours)
    }
}

