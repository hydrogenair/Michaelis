import SwiftUI
import Charts

struct SleepQualityView: View {
    @Environment(\.presentationMode) var presentationMode
    // 模拟数据，包含深睡、浅睡和清醒时长（以小时计）
    let weeklyData = [
        ("周一", [SleepStage(type: "深睡", hours: 5.0), SleepStage(type: "浅睡", hours: 2.0), SleepStage(type: "清醒", hours: 0.5)]),
        ("周二", [SleepStage(type: "深睡", hours: 4.5), SleepStage(type: "浅睡", hours: 2.5), SleepStage(type: "清醒", hours: 0.5)]),
        ("周三", [SleepStage(type: "深睡", hours: 5.0), SleepStage(type: "浅睡", hours: 1.5), SleepStage(type: "清醒", hours: 1.0)]),
        ("周四", [SleepStage(type: "深睡", hours: 3.5), SleepStage(type: "浅睡", hours: 3.0), SleepStage(type: "清醒", hours: 1.0)]),
        ("周五", [SleepStage(type: "深睡", hours: 5.5), SleepStage(type: "浅睡", hours: 2.0), SleepStage(type: "清醒", hours: 1.0)]),
        ("周六", [SleepStage(type: "深睡", hours: 6.0), SleepStage(type: "浅睡", hours: 1.5), SleepStage(type: "清醒", hours: 1.0)]),
        ("周日", [SleepStage(type: "深睡", hours: 5.0), SleepStage(type: "浅睡", hours: 2.0), SleepStage(type: "清醒", hours: 0.5)])
    ]
    
    var body: some View {
        ZStack {
            // 背景渐变色
            LinearGradient(
                gradient: Gradient(colors: [
                    Color(red: 255 / 255, green: 255 / 255, blue: 232 / 255),
                    Color(red: 217 / 255, green: 255 / 255, blue: 224 / 255)
                ]),
                startPoint: .top,
                endPoint: .bottom
            )
            .ignoresSafeArea()
            
            VStack(spacing:20){
                // 顶部导航栏
                HStack {
                    Button(action: {
                        presentationMode.wrappedValue.dismiss()
                    }) {
                        Image(systemName: "chevron.left")
                            .font(.title2)
                            .foregroundColor(.black)
                    }
                    Spacer()
                    Text("睡眠质量")
                        .font(.headline)
                    Spacer()
                    Image(systemName: "square.and.arrow.up")
                    
                }
                .padding(.horizontal)
                
             
                
                VStack{
                    // 日期选择器
                    Image("sleepy").frame(width: 30, height: 20 )
                    HStack {
                        Spacer()
                        Text("日")
                        Spacer()
                        Text("周")
                            .foregroundColor(.purple)
                            .fontWeight(.bold)
                        Spacer()
                        Text("月")
                        Spacer()
                    }
                    .padding(.vertical, 5)
                    .background(Color(.systemGray5))
                    .cornerRadius(8)
                    .padding(.horizontal)
                    
                    // 柱状图
                    Chart {
                        ForEach(weeklyData, id: \.0) { day, stages in
                            ForEach(stages) { stage in
                                BarMark(
                                    x: .value("Day", day),
                                    y: .value("Hours", stage.hours)
                                )
                                .foregroundStyle(by: .value("Type", stage.type))
                            }
                        }
                    }
                    .frame(height: 200)
                    .padding(.horizontal)
                    .chartForegroundStyleScale([
                        "深睡": .blue,
                        "浅睡": .purple,
                        "清醒": .cyan
                    ])
                    .chartYAxis {
                        AxisMarks(position: .leading)
                    }
                    .padding(.top, 10)
                    
                    // 本周睡眠统计
                    VStack(spacing: 8) {
                        HStack{
                            Text("本周睡眠")
                                .font(.headline)
                                .padding(.top)
                            Image("deep-sleep").frame(width: 30, height: 10 )
                        }
                        
                        
                        VStack(spacing: 8) {
                            SleepDataRow(title: "平均睡眠时长", value: "8小时30分")
                            SleepDataRow(title: "平均清醒", value: "8小时30分")
                            SleepDataRow(title: "平均浅睡", value: "8小时30分")
                            SleepDataRow(title: "平均深睡", value: "8小时30分")
                        }
                        .padding()
                        .background(Color.yellow.opacity(0.2))
                        .cornerRadius(10)
                    }
                    .padding(.horizontal)
                    
                    // 管家推荐
                    HStack {
                        HStack {
                            Text("管家推荐")
                                .font(.caption)
                                .foregroundColor(Color.white)
                                .padding(.horizontal)
                                .padding(.vertical, 4)
                                .background(Color.black)
                                .cornerRadius(10)
                            Image("chatroom")
                                .resizable()
                                .frame(width: 50, height: 50)
                        }
                        Spacer()
                        Button("更多") {
                            // 查看更多
                        }
                    }
                    .padding(.horizontal)
                    .padding(.top)
                    
                    // 推荐项目滑动视图
                    ScrollView(.horizontal, showsIndicators: false) {
                        HStack(spacing: 10) {
                            RecommendationView(title: "睡前音乐", description: "睡前聆听助眠音乐，放松心情快速入睡", imageName: "back1")
                            RecommendationView(title: "睡前晚安", description: "睡前道晚安，提醒按时睡觉，记录提醒事项", imageName: "back2")
                            RecommendationView(title: "助眠白噪音", description: "白噪音能够制造一个相对均匀的声音背景，可以抑制来自外部环境的干扰声音，减少对睡眠的干扰", imageName: "back3")
                        }
                        .padding(.horizontal)
                    }
                    .padding(.bottom)
                }
            } .padding(.top,-50)
        }
    }
}

struct SleepStage: Identifiable {
    let id = UUID()
    let type: String
    let hours: Double
}

struct SleepDataRow: View {
    let title: String
    let value: String
    
    var body: some View {
        HStack {
            Text(title)
                .font(.subheadline)
            Spacer()
            Text(value)
                .font(.subheadline)
        }
        .overlay(
            Rectangle()
                .frame(height: 1)
                .foregroundColor(Color.gray)
                .opacity(0.5),
            alignment: .bottom
        )
    }
}

struct RecommendationView: View {
    let title: String
    let description: String
    let imageName: String
    
    var body: some View {
        VStack(alignment: .leading, spacing: 5) {
            Text(title)
                .font(.headline)
                .foregroundColor(.white)
            Text(description)
                .font(.subheadline)
                .foregroundColor(.white)
            
        }
        .padding()
        .background(
            Image(imageName)
                .resizable()
                .scaledToFill()
        )
        .cornerRadius(10)
        .frame(width: 160, height: 120)
    }
}

struct SleepQualityView_Previews: PreviewProvider {
    static var previews: some View {
        SleepQualityView()
    }
}
