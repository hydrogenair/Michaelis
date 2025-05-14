import SwiftUI

struct HealthRecordsView: View {
    let healthRecords = [
        ("睡眠质量", "moon.zzz.fill", Color.purple),
        ("心率心跳", "heart.fill", Color.red),
        ("卡路里", "flame.fill", Color.orange),
        ("体重体脂", "scalemass.fill", Color.blue),
        ("体温", "thermometer.sun.fill", Color.yellow),
        ("血氧血糖", "drop.fill", Color.red)
    ]
    @Environment(\.presentationMode) var presentationMode

    var body: some View {
        VStack {
            HStack {
                Button(action: {
                    presentationMode.wrappedValue.dismiss()
                }) {
                    Image(systemName: "chevron.left")
                        .font(.title)
                }
                Spacer()
                Button(action: {
                    // More options button action
                }) {
                    Image(systemName: "ellipsis")
                        .font(.title)
                }
            }
            .padding()

            // Header Section
            VStack(alignment: .leading, spacing: 10) {
                Text("专属管家提醒：")
                    .font(.headline)
                    .padding(.top)
                Text("主人，压力大的时候，适当的放松和休息是非常重要的。或者已被温暖的香和一些轻柔佛如影月可以帮助您放松。请相信，您的舒适和健康始终是我的首要任务！")
                    .font(.subheadline)
                    .foregroundColor(.secondary)
            }
            .padding()
            .background(
                LinearGradient(gradient: Gradient(colors: [
                    Color(red: 255/255, green: 176/255, blue: 179/255, opacity: 1),
                    Color(red: 255/255, green: 226/255, blue: 222/255, opacity: 1),
                    Color(red: 250/255, green: 208/255, blue: 196/255, opacity: 1),
                    Color(red: 255/255, green: 237/255, blue: 232/255, opacity: 1)
                ]), startPoint: .topLeading, endPoint: .bottomTrailing)
            )
            .cornerRadius(30)
            .padding(.horizontal)

            // "向我倾诉" Component
            NavigationLink(destination: ServantView().navigationBarBackButtonHidden(true)) {
                HStack {
                    Spacer()
                    Text("向我倾诉")
                        .font(.caption)
                        .foregroundColor(Color.white)
                        .padding(.horizontal)
                        .padding(.vertical, 4)
                        .background(Color.black)
                        .cornerRadius(10)
                    Image("chatroom")
                        .resizable()
                        .frame(width: 40, height: 40)
                }
                .padding(.horizontal)
            }

            VStack {
                // Health Records List Section
                Text("Health records")
                    .font(.title2)
                    .bold()
                    .padding(.top)

                List {
                    ForEach(healthRecords, id: \.0) { record in
                        NavigationLink(destination: destinationView(for: record.0).navigationBarBackButtonHidden(true)) {
                            RecordRow(record: record)
                        }
                        .buttonStyle(PlainButtonStyle()) // 去掉导航自带的箭头
                    }
                    .listRowBackground(Color.clear)
                }
                .listStyle(PlainListStyle())
                .padding(.horizontal)
            }
            .background(
                LinearGradient(gradient: Gradient(colors: [
                    Color(red: 218/255, green: 250/255, blue: 215/255, opacity: 1),
                    Color(red: 245/255, green: 236/255, blue: 208/255, opacity: 1),
                    Color(red: 240/255, green: 247/255, blue: 220/255, opacity: 1),
                    Color(red: 242/255, green: 219/255, blue: 216/255, opacity: 1)
                ]), startPoint: .leading, endPoint: .trailing)
            )
        }
    }

    // Function to return the appropriate destination view based on record name
    @ViewBuilder
    func destinationView(for recordName: String) -> some View {
        switch recordName {
        case "睡眠质量":
            SleepQualityView()
        case "心率心跳":
            HeartRateView()
        case "卡路里":
            CalorieView()
        case "体重体脂":
            WeightView()
        case "体温":
            TemperatureView()
        case "血氧血糖":
            OxygenView()
        default:
            EmptyView()
        }
    }
}

struct RecordRow: View {
    let record: (String, String, Color)

    var body: some View {
        HStack {
            Image(systemName: record.1)
                .foregroundColor(record.2)
            Text(record.0)
            Spacer()
        }
        .padding()
        .background(
            LinearGradient(gradient: Gradient(colors: [
                Color.white.opacity(0.3),
                Color(red: 1.0, green: 0.96, blue: 0.96, opacity: 0.3),
                Color(red: 1.0, green: 0.86, blue: 0.86, opacity: 0.3),
                Color(red: 1.0, green: 0.71, blue: 0.71, opacity: 0.3)
            ]), startPoint: .leading, endPoint: .trailing)
        )
        .cornerRadius(10)
        .overlay(
            RoundedRectangle(cornerRadius: 10)
                .stroke(Color.white, lineWidth: 1)  // White border
        )
        .shadow(color: .gray, radius: 5, x: 0, y: 2)
    }
}

struct HealthRecordsView_Previews: PreviewProvider {
    static var previews: some View {
        HealthRecordsView()
    }
}
