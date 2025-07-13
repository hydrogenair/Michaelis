import SwiftUI
import Charts

struct CalorieView: View {
    @Environment(\.presentationMode) var presentationMode
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
                Text("卡路里详情")
                    .font(.headline)
                    .padding()
                Spacer()
                Button(action: {
                    // Share button action
                }) {
                    Image(systemName: "square.and.arrow.up")
                        .font(.title)
                        .padding(.trailing)
                }
            }
            
            ScrollView {
                VStack(alignment: .leading, spacing: 16) {
                    HStack {
                        Text("今日卡路里")
                            .font(.title2)
                            .padding(5.0)
                        Image("flame").resizable().frame(width: 20,height:20)
                    }
                    
                    // Bar Chart
                    BarChartView(title: "")
                        .frame(height: 200)
                        .padding(.horizontal)
                    
                    // Calorie Summary
                    HStack {
                        Text("今天你消耗了1308.4kcal, 相当于一块牛排的热量")
                            .font(.body)
                        
                    }
                    .padding()
                    .background(
                            LinearGradient(gradient: Gradient(colors: [
                                Color(red: 1, green: 0.69, blue: 0.70, opacity: 0.8),
                                Color(red: 0.98, green: 0.81, blue: 0.77, opacity: 0.8),
                                Color(red: 0.98, green: 0.81, blue: 0.77, opacity: 0.8)
                            ]), startPoint: .leading, endPoint: .trailing)
                        )
                        .cornerRadius(10)
                        .shadow(color: Color.black.opacity(0.25), radius: 4, x: 4, y: 2)
                    .padding(.horizontal)
                    
                    HStack {
                        Spacer()
                        Text("继续加油保持！")
                            .font(.body)
                            .padding()
                            .background(
                                LinearGradient(gradient: Gradient(colors: [
                                    Color(red: 1, green: 0.69, blue: 0.70, opacity: 0.8),
                                    Color(red: 0.98, green: 0.81, blue: 0.77, opacity: 0.8),
                                    Color(red: 0.98, green: 0.81, blue: 0.77, opacity: 0.8)
                                ]), startPoint: .leading, endPoint: .trailing)
                            )
                            .cornerRadius(10)
                            .shadow(color: Color.black.opacity(0.25), radius: 4, x: 4, y: 2)
                            .padding(.horizontal)
                    }
                    
                    // Total and Active Calorie
                    HStack(spacing: 16) {
                        VStack {
                            Text("Total")
                            Image("kaluli").resizable().frame(width: 40, height: 40)
                            Text("1308.4 kcal")
                        }
                        .frame(maxWidth: .infinity)
                        .padding()
                        .background(Color.white.opacity(0.75))
                        .cornerRadius(12)
                        .shadow(color: Color.black.opacity(0.25), radius: 4, x: 2, y: 2)
                        
                        VStack {
                            Text("Active calories")
                            Image("kaluli").resizable().frame(width: 40, height: 40)
                            Text("183 kcal")
                        }
                        .frame(maxWidth: .infinity)
                        .padding()
                        .background(Color.white.opacity(0.75))
                        .cornerRadius(12)
                        .shadow(color: Color.black.opacity(0.25), radius: 4, x: 2, y: 2)
                    }
                    .padding(.horizontal)

                    
                    // Today's Ranking
                    VStack(alignment: .leading, spacing: 8) {
                        HStack {
                            Image( "rank").resizable().frame(width: 30, height: 30)
                            Text("今日排名")
                                .font(.subheadline)
                            Text("up")
                                .foregroundColor(Color(red: 252/255, green: 138/255, blue: 88/255, opacity: 1))
                                .padding(.horizontal, 6)
                                .padding(.vertical, 2)
                                .background(Color(red: 1, green: 133/255, blue: 137/255, opacity: 0.2))
                                .cornerRadius(5)
                            Spacer()
                            Text("点击查看详情")
                                .font(.subheadline)
                                .foregroundColor(Color.gray)
                                .padding(.leading)
                            Image(systemName: "chevron.right")
                                .font(.title2)
                                .padding(.leading)
                            
                        }
                        
                        HStack {
                            Text("今天稍有松懈哦, 要相信自己的力量")
                                .padding(.horizontal)
                                .padding(.vertical, 4)
                            Spacer()
                        }
                        
                        
                    }
                    .padding()
                    .background(Color(red: 1, green: 176/255, blue: 179/255).opacity(0.4))
                    .cornerRadius(12)
                    .padding(.horizontal)
                }
            }
            
        }.background(
            LinearGradient(gradient: Gradient(colors: [Color(red: 1, green: 1, blue: 0.91), Color(red: 0.85, green: 1, blue: 0.88)]), startPoint: .top, endPoint: .bottom)
                .edgesIgnoringSafeArea(.all)
        )
    }
}


struct BarChartView: View {
    @State private var animatedData: [(String, Double)] = [
        ("00:00-8:00", 0),
        ("8:00-12:00", 0),
        ("12:00-18:00", 0),
        ("18:00-24:00", 0)
    ]
    
    var data: [(String, Double)] = [
        ("00:00-8:00", 4),
        ("8:00-12:00", 8),
        ("12:00-18:00", 5),
        ("18:00-24:00", 12)
    ]
    
    var title: String

    var body: some View {
        VStack(alignment: .leading) {
            Text(title)
                .font(.headline)
            
            Chart {
                ForEach(animatedData, id: \.0) { entry in
                    BarMark(
                        x: .value("Time", entry.0),
                        y: .value("Calories", entry.1)
                    )
                    .foregroundStyle(LinearGradient(
                        gradient: Gradient(colors: [
                            Color(red: 255/255, green: 154/255, blue: 158/255, opacity: 1),
                            Color(red: 250/255, green: 208/255, blue: 196/255, opacity: 1)
                        ]),
                        startPoint: .top,
                        endPoint: .bottom
                    ))
                }
            }
            .chartXAxis {
                AxisMarks(values: data.map { $0.0 }) { value in
                    AxisValueLabel {
                        Text(value.as(String.self) ?? "")
                    }
                }
            }
            .chartYAxis {
                AxisMarks(values: .stride(by: 2))
            }
        }
        .padding()
        .onAppear {
            animateChart()
        }
    }
    
    func animateChart() {
        for (index, value) in data.enumerated() {
            DispatchQueue.main.asyncAfter(deadline: .now() + Double(index) * 0.3) {
                withAnimation(.easeInOut(duration: 0.5)) {
                    animatedData[index].1 = value.1
                }
            }
        }
    }
}

struct CalorieView_Previews: PreviewProvider {
    static var previews: some View {
        CalorieView()
    }
}
