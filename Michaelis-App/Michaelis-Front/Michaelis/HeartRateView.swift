import SwiftUI

struct HeartRateView: View {
    
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
            VStack(spacing:10) {
                
                TopNavigationBar()
                
                HeartRateDisplay()
                
                LineChartView(data: [85, 70, 102, 112,80],
                              timeLabels: ["00:00", "06:00", "12:00", "18:00", "24:00"],
                              heartRateLabels: [120, 110, 100, 90,80,70, 60])
                    .frame(height: 200)
                    .padding()
                
                HeartRateMinMax().padding(.top)
                
                ActivityReminder()
                
                Spacer()
            }
            .padding(.top,-20)
        }
        }
}

struct TopNavigationBar: View {
    @Environment(\.presentationMode) var presentationMode
    var body: some View {
        HStack {
            Button(action: {
                presentationMode.wrappedValue.dismiss()
            }) {
                Image(systemName: "arrow.backward")
                    .foregroundColor(Color.black)
            }
            Spacer()
            Text("心率心跳")
                .font(.title2)
            Spacer()
            Button(action: {
                // 分享操作
            }) {
                Image(systemName: "square.and.arrow.up")
            }
            Button(action: {
                // 更多操作
            }) {
                Image(systemName: "ellipsis")
            }
        }
        .padding()
    }
}

struct HeartRateDisplay: View {
    var body: some View {
        VStack(spacing: 8) {
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
            HStack{
                Image("rate")
                HStack{
                    Text("90")
                        .font(.title2).foregroundColor(Color.pink).padding(.trailing, -5)
                        
                    Text("次/min")
                        .font(.footnote)
                        .foregroundColor(Color.pink)
                        .padding(.bottom, -4.0)
                    
                }
                Spacer()
                
            }.padding(.leading)
            
        }
    }
}

struct HeartRateMinMax: View {
    var body: some View {
        HStack {
            Spacer()
            VStack {
                Text("最高")
                    .font(.subheadline)
                Image("rate")
                    .foregroundColor(.red)
                HStack{
                    Text("125")
                        .font(.title2).padding(.trailing, -5)
                    Text("次/min")
                        .font(.footnote)
                        .padding(.bottom, -4.0)
                    
                }
                
            }.padding(.vertical, 5)
                .padding(.horizontal)
                .background(.pink.opacity(0.2))
                .cornerRadius(8)
                
            Spacer()
            VStack {
                Text("最低")
                    .font(.subheadline)
                Image("rate")
                    .foregroundColor(.red)
                HStack{
                    Text(" 70")
                        .font(.title2).padding(.trailing, -5)
                    Text("次/min")
                        .font(.footnote)
                        .padding(.bottom, -4.0)
                    
                }
                
            }.padding(.vertical, 5)
                .padding(.horizontal)
                .background(.pink.opacity(0.2))
                .cornerRadius(8)
            Spacer()
            
            
        }
        .padding()
    }
}

struct ActivityReminder: View {
    var body: some View {
        VStack(alignment: .leading, spacing: 8) {
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
            }
            
            ProgressBarView(progress: 0.7)
                .frame(height: 20)
            
            // 图例
            HStack {
                HStack {
                    Circle().frame(width: 10, height: 10).foregroundColor(.green)
                    Text("24分钟\n热身")
                }
                Spacer()
                HStack {
                    Circle().frame(width: 10, height: 10).foregroundColor(.blue)
                    Text("7分钟\n燃脂")
                }
                Spacer()
                HStack {
                    Circle().frame(width: 10, height: 10).foregroundColor(.yellow)
                    Text("0分钟\n有氧")
                }
                Spacer()
                HStack {
                    Circle().frame(width: 10, height: 10).foregroundColor(.orange)
                    Text("0分钟\n无氧")
                }
            }
            .font(.footnote)
        }
        .padding()
    }
}

struct ProgressBarView: View {
    var progress: CGFloat
    
    var body: some View {
        ZStack(alignment: .leading) {
            Rectangle()
                .fill(Color.gray.opacity(0.3))
                .cornerRadius(10)
            Rectangle()
                .fill(Color.green)
                .frame(width: progress * 300) // 根据进度调整宽度
                .cornerRadius(10)
            Rectangle()
                .fill(Color.blue)
                .frame(width: progress * 50) // 根据进度调整宽度
                .cornerRadius(10)
        }
    }
}

struct LineChartView: View {
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
                                .offset(x: -192).offset(y:8) // 左移到边框外
                            
                            // X轴标签（放在边框外）
                            createXAxisLabels(in: geometry)
                                .offset(x:-5,y: 115)
                                // 下移到边框外
                        }
                        .background(Color.mint.opacity(0.1).cornerRadius(10))
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
        let maxY = data.max() ?? 1
        let minY = data.min() ?? 0
        let yRange = 60
        
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
        let minValue = data.min() ?? 0
        let maxValue = data.max() ?? 1
        let range = 60
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
                    Text(String(format: "%.0f", heartRateLabels[index]))

                        .font(.caption)
                        .foregroundColor(.gray)
                        .frame(height: 22, alignment: index == 0 ? .top : .center)
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

struct HeartRateView_Previews: PreviewProvider {
    static var previews: some View {
        HeartRateView()
    }
}
