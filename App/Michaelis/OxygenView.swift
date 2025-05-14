import SwiftUI

struct OxygenView: View {
    @Environment(\.presentationMode) var presentationMode
    var body: some View {
        ZStack{
            LinearGradient(
                gradient: Gradient(colors: [
                    Color(red: 255 / 255, green: 255 / 255, blue: 232 / 255),
                    Color(red: 217 / 255, green: 255 / 255, blue: 224 / 255)
                ]),
                startPoint: .top,
                endPoint: .bottom
            )
            .ignoresSafeArea()
            VStack(spacing: 16) {
                // 顶部导航栏
                
                HStack {
                    Button(action: {  presentationMode.wrappedValue.dismiss() }) {
                        Image(systemName: "arrow.left")
                            .foregroundColor(Color.black)
                    }
                    Spacer()
                    Text("血氧")
                        .font(.title2)
                        .bold()
                    Spacer()
                    HStack(spacing: 16) {
                        Image(systemName: "square.and.arrow.up")
                        Image(systemName: "ellipsis")
                    }
                }
                .padding(.horizontal)
                
                // 日期选择栏
                HStack {
                    Spacer()
                    Text("日")
                        .foregroundColor(.pink)
                        .fontWeight(.bold)
                    Spacer()
                    Text("周")
                    Spacer()
                    Text("月")
                    Spacer()
                }
                .padding(.vertical, 5)
                .background(.pink.opacity(0.2))
                .cornerRadius(8)
                .padding(.horizontal)
                
                // 血氧数值
                HStack {
                    Image("blood1")
                        .foregroundColor(.red)
                        .frame(width: 20,height: 10)
                    Text("98.8%")
                        .font(.title2)
                        .bold()
                    Spacer()
                }
        
                
                // 折线图
                LineChartView2(
                    data: [80, 100, 115,113,99],
                    timeLabels: ["00:00", "06:00", "12:00", "18:00", "00:00"],
                    heartRateLabels: [100, 99, 98, 97, 96, 95]
                ).frame(width: 350,height: 200)
               
                
                // 最低和最高血氧卡片
                HStack(spacing: 16) {
                    VStack {
                        Text("最低")
                        Image( "blood2")
                            .foregroundColor(.red)
                        Text("96.2%")
                            .font(.title2)
                            .bold()
                    }
                    .frame(maxWidth: .infinity)
                    .padding()
                    .background(Color(red: 255 / 255, green: 255 / 255, blue: 232 / 255).opacity(0.9))
                    
                    VStack {
                        Text("最高")
                        Image("blood2")
                            .foregroundColor(.red)
                        Text("99.5%")
                            .font(.title2)
                            .bold()
                    }
                    .frame(maxWidth: .infinity)
                    .padding()
                    .background(Color(red: 255 / 255, green: 255 / 255, blue: 232 / 255).opacity(0.9))
                }.padding(.top)
                
                // 管家提醒和血氧较低提醒
                HStack(spacing: 16) {
                    HStack (spacing: 5){
                        Image("chatroom")
                            .resizable()
                            .frame(width: 50, height: 50)
                        Text("管家推荐")
                            .font(.caption)
                            .foregroundColor(Color.white)
                            .padding(.horizontal)
                            .padding(.vertical, 4)
                            .background(Color.black)
                            .cornerRadius(10)
                        
                    }
                    
                    
                    Text("以下时间血氧较低哦")
                        .font(.subheadline)
                        .foregroundColor(.pink)
                        .padding(.horizontal)
                        .background(RoundedRectangle(cornerRadius: 10).fill(Color.pink.opacity(0.2)))
                }
                .padding(.vertical, 8)
                
                // 血氧记录列表
                VStack(spacing: 8) {
                    HStack {
                        Text("95%")
                            .font(.headline)
                        Spacer()
                        Text("x月x日 14:21")
                        Image(systemName: "chevron.right")
                            .foregroundColor(Color.white)
                    }
                    .padding()
                    .background(RoundedRectangle(cornerRadius: 10).fill(Color(red: 247 / 255, green: 217 / 255, blue: 203 / 255).opacity(0.9))).padding(.horizontal)
                    
                    HStack {
                        Text("95%")
                            .font(.headline)
                        Spacer()
                        Text("x月x日 14:21")
                        Image(systemName: "chevron.right")
                            .foregroundColor(Color.white)
                    }
                    .padding()
                    .background(RoundedRectangle(cornerRadius: 10).fill(Color(red: 247 / 255, green: 217 / 255, blue: 203 / 255).opacity(0.9)))
                    .padding(.horizontal)
                }
                Spacer()
            }
            .padding(.top,-5)
        }
        }
        
}

struct LineChartView2: View {
    var data: [Double]
    var timeLabels: [String]
    var heartRateLabels: [Double]
    
    var body: some View {
        GeometryReader { geometry in
                    VStack {
                        ZStack {
                            // 绘制网格背景
                            createGrid(in: geometry)
                            
                            // 折线图路径
                            createPath(in: geometry)
                                .stroke(Color.red, lineWidth: 2)
                                .overlay(
                                    createDataPoints(in: geometry)
                                )
                            
                            // Y轴标签（放在边框外）
                            createYAxisLabels(in: geometry)
                                .offset(x: -194).offset(y:8) // 左移到边框外
                            
                            // X轴标签（放在边框外）
                            createXAxisLabels(in: geometry)
                                .offset(x:-5,y: 110)
                                // 下移到边框外
                        }
                        .background(Color.white.opacity(0.7).cornerRadius(10))
                        .offset(x:8)
                    }
                }
    }
    // 创建网格背景
        private func createGrid(in geometry: GeometryProxy) -> some View {
            let lineCount = heartRateLabels.count-1
            let lineSpacing = geometry.size.height / CGFloat(lineCount - 1)
            
            return VStack(spacing: lineSpacing) {
                ForEach(0..<lineCount, id: \.self) { _ in
                    Rectangle()
                        .fill(Color.green.opacity(0.3))
                        .frame(height: 1)
                        .overlay(Divider().frame(height: 1).background(Color.gray.opacity(0.3)))
                }
            }
            .frame(maxHeight: .infinity, alignment: .top)
        }
    
    private func createPath(in geometry: GeometryProxy) -> Path {
        var path = Path()
        
        guard data.count > 1 else { return path }
        
        let xStep = geometry.size.width / CGFloat(data.count - 1)
        _ = data.max() ?? 1
       _ = data.min() ?? 0
        _ = 60
        
        var previousPoint = CGPoint(
            x: 0,
            y: geometry.size.height * CGFloat(1 - (data[0] - 60) / 60)
        )
        
        path.move(to: previousPoint)
        
        for index in 1..<data.count {
            let xPosition = CGFloat(index) * xStep
            let yPosition = geometry.size.height * CGFloat(1 - (data[index] - 60) / 60)
            let currentPoint = CGPoint(x: xPosition, y: yPosition)
            
            let midPoint = CGPoint(
                x: (previousPoint.x + currentPoint.x) / 2,
                y: (previousPoint.y + currentPoint.y) / 2
            )
            
            path.addQuadCurve(to: midPoint, control: CGPoint(
                x: previousPoint.x,
                y: midPoint.y
            ))
            
            path.addQuadCurve(to: currentPoint, control: CGPoint(
                x: currentPoint.x,
                y: midPoint.y
            ))
            
            previousPoint = currentPoint
        }
        
        return path
    }
    
    private func dataPointPosition(index: Int, geometry: GeometryProxy) -> CGPoint {
        let width = geometry.size.width
        let height = geometry.size.height
        
        // Separate calculations for xPosition
        let count = CGFloat(data.count - 1)
        let xMultiplier = CGFloat(index)
        let xPosition: CGFloat = (width / count) * xMultiplier
        
        // Separate calculations for yPosition
        _ = data.min() ?? 0
        _ = data.max() ?? 1
        _ = 60
        let normalizedValue = (data[index] - 60) / 60
        let yPosition: CGFloat = height * CGFloat(1 - normalizedValue)
        
        return CGPoint(x: xPosition, y: yPosition)
    }


    
    private func createDataPoints(in geometry: GeometryProxy) -> some View {
        ForEach(0..<data.count, id: \.self) { index in
            Circle()
                .fill(Color.red)
                .frame(width: 6, height: 6)
                .position(dataPointPosition(index: index, geometry: geometry))
        }
    }
    
    private func createYAxisLabels(in geometry: GeometryProxy) -> some View {
        _ = geometry.size.height / CGFloat(heartRateLabels.count - 1)
            
            return VStack {
                ForEach(0..<heartRateLabels.count, id: \.self) { index in
                    Text(String(format: "%.f", heartRateLabels[index]))

                        .font(.caption)
                        .foregroundColor(.gray)
                        .frame(height: 25, alignment: index == 0 ? .top : .center)
                }
            }
        }
        
        
    
    private func createXAxisLabels(in geometry: GeometryProxy) -> some View {
        _ = geometry.size.width / CGFloat(timeLabels.count - 1)
            
            return HStack {
                ForEach(0..<timeLabels.count, id: \.self) { index in
                    Text(timeLabels[index])
                        .font(.caption)
                        .foregroundColor(.gray)
                        .frame(width: 65, alignment: index == 0 ? .leading : .center)
                }
            }
        }
}


struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        OxygenView()
    }
}

