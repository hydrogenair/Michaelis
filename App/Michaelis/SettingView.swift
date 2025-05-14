import SwiftUI

struct ServiceButton: View {
    var imageName: String
    var title: String
    
    var body: some View {
        HStack {
            Image(imageName)
                .resizable()
                .frame(width: 30, height: 30)
                .padding(.leading, 20)
            Text(title)
                .font(.system(size: 18))
                .foregroundColor(Color.black)
                .padding(.leading, 10)
            Spacer()
            Image(systemName: "chevron.right")
                .foregroundColor(Color.black)
                .padding(.trailing, 20)
        }
        .frame(width: UIScreen.main.bounds.width * 0.9, height: 60)
        .background(
            LinearGradient(gradient: Gradient(colors: [
                Color.white.opacity(0.3),
                Color(red: 1.0, green: 0.96, blue: 0.96, opacity: 0.3),
                Color(red: 1.0, green: 0.86, blue: 0.86, opacity: 0.3),
                Color(red: 1.0, green: 0.71, blue: 0.71, opacity: 0.3)
            ]), startPoint: .leading, endPoint: .trailing)
        )
        .cornerRadius(10)
        .shadow(color: Color.gray.opacity(0.7), radius: 5, x: 0, y: 2)
        .padding(.horizontal, UIScreen.main.bounds.width * 0.05)
    }
}

struct UpgradeServiceButton: View {
    var imageName: String
    var title: String
    
    var body: some View {
        VStack {
            Image(imageName)
                .resizable()
                .frame(width: 90, height: 60)
            Text(title)
                .font(.system(size: 16))
        }
        .frame(width: 100, height: 120)
        .foregroundColor(Color.black)
        .background(Color.white)
        .cornerRadius(10)
        .shadow(color: Color.gray.opacity(0.2), radius: 5, x: 0, y: 2)
        .padding(.horizontal, UIScreen.main.bounds.width * 0.05)
    }
}

struct SettingView: View {
    @Environment(\.presentationMode) var presentationMode
    @State private var showCustomServantList = false
    
    var body: some View {
        ZStack {
            LinearGradient(gradient: Gradient(colors: [Color(red: 255/255, green: 255/255, blue: 232/255), Color(red: 217/255, green: 255/255, blue: 224/255)]), startPoint: .top, endPoint: .bottom)
                .edgesIgnoringSafeArea(.all)
            
            VStack(alignment: .leading, spacing: 20) {
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
                        Image(systemName: "ellipsis")
                            .font(.title)
                            .padding(.trailing)
                    }
                }
                .padding(.top, 10)
                
                ScrollView {
                    Group {
                        Text("专属服务")
                            .font(.headline)
                            .padding(.leading, 20)
                        
                        Button(action: {
                            withAnimation {
                                showCustomServantList.toggle()
                            }
                        }) {
                            ServiceButton(imageName: "servant", title: "定制管家形象")
                        }
                        
                        if showCustomServantList {
                            VStack(spacing: 10) {
                                CustomServantRow(imageName: "wizardAvatar", title: "魔法小猫")
                                CustomServantRow(imageName: "duola", title: "哆啦A梦")
                                CustomServantRow(imageName: "ymhd", title: "樱木花道")
                                CustomServantRow(imageName: "ldy", title: "林黛玉")
                                CustomServantRow(imageName: "", title: "自定义形象...")
                            }
                            .transition(.slide)
                        }
                        
                        Text("基础服务")
                            .font(.headline)
                            .padding(.leading, 20)
                        
                        ServiceButton(imageName: "emotionraft", title: "情绪评估")
                        ServiceButton(imageName: "healthmanage", title: "健康管理")
                        ServiceButton(imageName: "privacy", title: "隐私设置")
                        
                        HStack {
                            Image("star")
                                .resizable()
                                .frame(width: 20, height: 20)
                            Text("升级服务")
                                .font(.headline)
                        }
                        .padding(.leading, 20)
                    }
                    
                    VStack {
                        HStack {
                            Spacer()
                            NavigationLink(destination: EmotionWindowView().navigationBarBackButtonHidden(true)) {
                                                                UpgradeServiceButton(imageName: "emotinservice", title: "情绪服务")
                                                            }
                            UpgradeServiceButton(imageName: "design", title: "个性化计划")
                            Spacer()
                        }
                        HStack {
                            Spacer()
                            UpgradeServiceButton(imageName: "report", title: "健康报告")
                            UpgradeServiceButton(imageName: "emergency", title: "紧急情况")
                            Spacer()
                        }
                    }
                    .padding(.horizontal)
                }
                
                Spacer()
            }
            .padding(.top, 10)
        }
    }
}

struct CustomServantRow: View {
    var imageName: String
    var title: String
    
    var body: some View {
        HStack {
            if !imageName.isEmpty {
                Image(imageName)
                    .resizable()
                    .frame(width: 30, height: 30)
                    .padding(.leading, 20)
            }
            Text(title)
                .font(.system(size: 18))
                .padding(.leading, 10)
            Spacer()
        }
        .frame(height: 60)
        .background(
            LinearGradient(gradient: Gradient(colors: [
                Color.white.opacity(0.3),
                Color(red: 1.0, green: 0.96, blue: 0.96, opacity: 0.3),
                Color(red: 1.0, green: 0.86, blue: 0.86, opacity: 0.3),
                Color(red: 1.0, green: 0.71, blue: 0.71, opacity: 0.3)
            ]), startPoint: .leading, endPoint: .trailing)
        )
        .cornerRadius(10)
        .shadow(color: Color.gray.opacity(0.7), radius: 5, x: 0, y: 2)
        .padding(.horizontal, 20)
    }
}

struct SettingView_Previews: PreviewProvider {
    static var previews: some View {
        SettingView()
    }
}
