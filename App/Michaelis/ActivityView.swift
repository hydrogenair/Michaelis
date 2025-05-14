import SwiftUI
import HealthKit
let healthStore = HKHealthStore()
func requestHealthKitAuthorization() {
    if HKHealthStore.isHealthDataAvailable() {
        let stepCountType = HKQuantityType.quantityType(forIdentifier: .stepCount)!
        let activeEnergyType = HKQuantityType.quantityType(forIdentifier: .activeEnergyBurned)!
        
        let dataTypesToRead: Set<HKObjectType> = [stepCountType, activeEnergyType]
        
        healthStore.requestAuthorization(toShare: nil, read: dataTypesToRead) { success, error in
            if !success {
                print("HealthKit authorization failed: \(String(describing: error))")
            }
        }
    }
}
func fetchStepCount(completion: @escaping (Double) -> Void) {
    let stepType = HKQuantityType.quantityType(forIdentifier: .stepCount)!
    let now = Date()
    let startOfDay = Calendar.current.startOfDay(for: now)
    
    let predicate = HKQuery.predicateForSamples(withStart: startOfDay, end: now, options: .strictStartDate)
    
    let query = HKStatisticsQuery(quantityType: stepType, quantitySamplePredicate: predicate, options: .cumulativeSum) { _, result, error in
        var steps = 0.0
        
        if let result = result, let sum = result.sumQuantity() {
            steps = sum.doubleValue(for: HKUnit.count())
        }
        
        DispatchQueue.main.async {
            completion(steps)
        }
    }
    
    healthStore.execute(query)
}

func fetchCaloriesBurned(completion: @escaping (Double) -> Void) {
    let calorieType = HKQuantityType.quantityType(forIdentifier: .activeEnergyBurned)!
    let now = Date()
    let startOfDay = Calendar.current.startOfDay(for: now)
    
    let predicate = HKQuery.predicateForSamples(withStart: startOfDay, end: now, options: .strictStartDate)
    
    let query = HKStatisticsQuery(quantityType: calorieType, quantitySamplePredicate: predicate, options: .cumulativeSum) { _, result, error in
        var calories = 0.0
        
        if let result = result, let sum = result.sumQuantity() {
            calories = sum.doubleValue(for: HKUnit.kilocalorie())
        }
        
        DispatchQueue.main.async {
            completion(calories)
        }
    }
    
    healthStore.execute(query)
}


struct ActivityView: View {
    @State private var showDatePicker = false
    @State private var selectedDate = Date()
    @State private var selectedActivity = "骑行"
    @State private var stepCount: Double = 0.0
    @State private var caloriesBurned: Double = 0.0
    @State private var sportHours: Double = 0.0
    let activities = ["跑步", "竞走", "游泳", "骑行"]
    @Environment(\.presentationMode) var presentationMode
    var body: some View {
        GeometryReader { geometry in
            ZStack {
                // Background Image
                Image("ActivityBackground")
                    .resizable()
                    .aspectRatio(contentMode: .fill)
                    .frame(width: geometry.size.width)

                
                VStack {
                    HStack {
                        Button(action: {
                            presentationMode.wrappedValue.dismiss()
                        }) {
                            Image(systemName: "chevron.left")
                                .font(.title)
                        }
                        Spacer()
                        Text("Hi，让我们一起运动吧!")
                            .font(.headline)
                        Spacer()
                        Button(action: {
                            // More options button action
                        }) {
                            Image(systemName: "ellipsis")
                                .font(.title)
                        }
                    }
                    .padding()
                    
                    // Date Buttons
                    HStack {
                        Button(action: {
                            // Today button action
                        }) {
                            Text("今天")
                                .font(.subheadline)
                                .padding()
                                .frame(width: 80, height: 40)
                                .background(Color.black)
                                .foregroundColor(.white)
                                .cornerRadius(10)
                        }
            
                        Spacer()
                        Button(action: {
                            showDatePicker.toggle()
                        }) {
                            Text("日期")
                                .font(.subheadline)
                                .padding()
                                .frame(width: 80, height: 40)
                                .background(Color.black)
                                .foregroundColor(.white)
                                .cornerRadius(10)
                        }
                    }
                    .padding()
                    if showDatePicker {
                                            DatePicker("选择日期", selection: $selectedDate, displayedComponents: .date)
                                                .datePickerStyle(GraphicalDatePickerStyle())
                                                .background(Color.white)
                                                .cornerRadius(10)
                                                .padding()
                                        }
                    
                    // Statistics
                    HStack(alignment: .top){
                        VStack {
                                       Text("步 数")
                                           .font(.headline)
                                       Image("foot")
                                           .resizable()
                                           .frame(width: 100, height: 100)
                                       Text("\(Int(stepCount))")
                                           .font(.largeTitle)
                                       Text("步")
                                           .font(.subheadline)
                                   }
                        .padding()
                                   .background(
                                       VisualEffectBlur(blurStyle: .light)
                                           .background(
                                               RoundedRectangle(cornerRadius: 36)
                                                   .fill(Color(red: 255/255, green: 181/255, blue: 182/255).opacity(0.3))
                                           )
                                           .clipShape(RoundedRectangle(cornerRadius: 36))
                                   )
                                   .frame(width: 180, height: 240)
                                   .shadow(radius: 5)

                                   // 卡路里和运动时长组件
                        VStack {
                                       VStack {
                                           Text("卡路里")
                                               .font(.headline)
                                               .padding(-4.0)
                                           Image("kaluli")
                                               .resizable()
                                               .frame(width: 40, height: 40)
                                           Text("\(Int(caloriesBurned))")
                                               .font(.title2)
                                               .padding(-4.0)
                                           Text("Calories")
                                               .font(.subheadline)
                            
                                       }
                                        
                                       .frame(width: 120, height: 120)
                                       .background(
                                        VisualEffectBlur(blurStyle: .light)
                                            .background(
                                                RoundedRectangle(cornerRadius: 36)
                                                    .fill(Color(red: 255/255, green: 181/255, blue: 182/255).opacity(0.3))
                                            )
                                            .clipShape(RoundedRectangle(cornerRadius: 20))
                                    )
                                       
                                       .shadow(radius: 5)

                                       VStack {
                                           Text("运动时长")
                                               .font(.headline)
                                               .padding(-4.0)
                                           Image("hour")
                                               .resizable()
                                               .frame(width: 40, height: 40)
                                           Text("\(String(format: "%.1f", sportHours))")
                                               .font(.title2)
                                               .padding(-4.0)
                                           Text("分钟")
                                               .font(.subheadline)
                                       }
                           .frame(width: 120, height: 120)
                           .background(
                            VisualEffectBlur(blurStyle: .light)
                                .background(
                                    RoundedRectangle(cornerRadius: 36)
                                        .fill(Color(red: 255/255, green: 181/255, blue: 182/255).opacity(0.3))
                                )
                                .clipShape(RoundedRectangle(cornerRadius: 20))
                        )
                                       .shadow(radius: 5)
                        }
                        .padding(0)
                    }
                                   .padding()
                        
                    // Achievement and Plan
                    VStack {
                        VStack {
                            Button(action: {
                                // Today's achievement action
                            }) {
                                HStack {
                                    Image("achievement")
                                        .resizable()
                                        .frame(width: 20, height: 20)
                                    Text("今日成就")
                                        .font(.body)
                                        .foregroundColor(Color.black)
                                        .padding(.leading, 10)
                                }
                                .padding()
                                .background(
                                    VisualEffectBlur(blurStyle: .light)
                                        .background(
                                            RoundedRectangle(cornerRadius: 36)
                                                .fill(Color(red: 255/255, green: 181/255, blue: 182/255).opacity(0.3))
                                        )
                                        .clipShape(RoundedRectangle(cornerRadius: 36))
                                )
                                .shadow(radius: 5)
                            }
                            .frame(maxWidth: .infinity)
                            .padding(.horizontal)
                            
                            Button(action: {
                                // My plan action
                            }) {
                                HStack {
                                    Image("plan")
                                        .resizable()
                                        .frame(width: 20, height: 20)
                                    Text("我的计划")
                                        .font(.body)
                                        .foregroundColor(Color.black)
                                        .padding(.leading, 10)
                                }
                                .padding()
                                .background(
                                    VisualEffectBlur(blurStyle: .light)
                                        .background(
                                            RoundedRectangle(cornerRadius: 36)
                                                .fill(Color(red: 255/255, green: 181/255, blue: 182/255).opacity(0.3))
                                        )
                                        .clipShape(RoundedRectangle(cornerRadius: 36))
                                )
                                .shadow(radius: 5)
                            }
                            .frame(maxWidth: .infinity)
                            .padding(.horizontal)
                        }
                        
                        // More Exercise
                        Button(action: {
                            // More exercise action
                        }) {
                            HStack {
                                Image("moresports")
                                    .resizable()
                                    .frame(width: 20, height: 20)
                                Text("更多运动")
                                    .font(.body)
                                    .foregroundColor(Color.black)
                                    .padding(.leading, 10)
                            }
                            .padding()
                            .background(
                                VisualEffectBlur(blurStyle: .light)
                                    .background(
                                        RoundedRectangle(cornerRadius: 36)
                                            .fill(Color(red: 255/255, green: 181/255, blue: 182/255).opacity(0.3))
                                    )
                                    .clipShape(RoundedRectangle(cornerRadius: 36))
                            )
                            .shadow(radius: 5)
                        }
                        .frame(maxWidth: .infinity)
                        .padding()
                    }
                    // Start Button
                    Spacer()
                    VStack {
                        Picker(selection: $selectedActivity, label: Text("")) {
                                            ForEach(activities, id: \.self) { activity in
                                                Text(activity)
                                                    .font(.largeTitle) // 更改字体大小
                                                    .foregroundColor(selectedActivity == activity ? Color.white : Color.black) // 更改选中的颜色
                                            }
                                        }
                                    .pickerStyle(MenuPickerStyle())
                                    .frame(width: 150, height: 50)
                                    .foregroundColor(.white)
                                    .cornerRadius(10)
                                    .accentColor(.white)
                                    .overlay(
                                        HStack {
                                            Image(systemName: "figure.run")
                                                .frame(width: 40, height: 40)
                                            Spacer()
                                        }
                                        .padding(.horizontal)
                                        .foregroundColor(.white)
                                    )
                                    
                        NavigationLink(destination: TimerView(sportHours: $sportHours).navigationBarBackButtonHidden(true)) {
                                        Text("START")
                                            .font(.title)
                                            .padding()
                                            .background(
                                                LinearGradient(gradient: Gradient(colors: [Color(red: 255/255, green: 102/255, blue: 102/255), Color(red: 255/255, green: 178/255, blue: 102/255)]), startPoint: .leading, endPoint: .trailing)
                                            )
                                            .foregroundColor(.white)
                                            .cornerRadius(10)
                                    }
                                }
                                .padding()
                                .background(
                                    Image("back")
                                        .scaleEffect(CGSize(width: 0.5, height: 0.2))
                                                            .frame(width: UIScreen.main.bounds.width * 2, height: UIScreen.main.bounds.height)
                                                            .clipped()
                                )
                                .padding()
                            }
                
            }
        }.background(
            LinearGradient(
                gradient: Gradient(colors: [
                    Color(red: 255 / 255, green: 255 / 255, blue: 232 / 255, opacity: 1),
                    Color(red: 217 / 255, green: 255 / 255, blue: 224 / 255, opacity: 1)
                ]),
                startPoint: .top,
                endPoint: .bottom
            )
        )
    }
    }

struct VisualEffectBlur: UIViewRepresentable {
        var blurStyle: UIBlurEffect.Style
        func makeUIView(context: Context) -> UIVisualEffectView {
            return UIVisualEffectView(effect: UIBlurEffect(style: blurStyle))
        }
        func updateUIView(_ uiView: UIVisualEffectView, context: Context) {}
    }
    

struct ActivityView_Previews: PreviewProvider {
    static var previews: some View {
        ActivityView()
    }
}
