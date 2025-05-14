import SwiftUI

struct LoginView: View {
    @State private var email: String = ""
    @State private var password: String = ""
    @State private var isAgreementChecked: Bool = false
    @State private var showAgreementAlert: Bool = false
    @State private var showLoginErrorAlert: Bool = false
    @State private var isLoggedIn: Bool = false  // 用于跟踪登录状态

    var body: some View {
        NavigationView {
            VStack {
                Spacer()
                
                Text("Hello!")
                    .font(.largeTitle)
                    .foregroundColor(.gray)
                    .padding(.bottom, 20)
                
                Image("login")
                    .resizable()
                    .frame(width: 80, height: 80)
                    .padding(.bottom, 30)
                
                HStack {
                    Image(systemName: "envelope")
                        .foregroundColor(.gray)
                        .padding(.horizontal, 7)
                    TextField("请输入您的邮箱……", text: $email)
                        .padding()
                    Spacer()
                }
                .background(Color.white.opacity(0.8))
                .cornerRadius(10)
                .padding(.horizontal, 20)
                
                HStack {
                    Image(systemName: "lock")
                        .foregroundColor(.gray)
                        .padding(.horizontal, 10)
                    SecureField("请输入您的密码……", text: $password)
                        .padding()
                }
                .background(Color.white.opacity(0.8))
                .cornerRadius(10)
                .padding(.horizontal, 20)
                
                HStack {
                    Spacer()
                    Text("忘记密码？")
                        .font(.footnote)
                        .foregroundColor(.gray)
                        .padding(.trailing, 16)
                }
                .padding(.bottom, 20)
                
                HStack {
                    NavigationLink(destination: RegisterView()) {
                        Text("注册账号")
                            .foregroundColor(.white)
                            .padding()
                            .frame(maxWidth: .infinity)
                            .background(Color.gray.opacity(0.5))
                            .cornerRadius(10)
                    }
                    
                    Button(action: {
                        if isAgreementChecked {
                            login()
                        } else {
                            showAgreementAlert = true
                        }
                    }) {
                        Text("立即登录")
                            .foregroundColor(.white)
                            .padding()
                            .frame(maxWidth: .infinity)
                            .background(Color.pink.opacity(0.7))
                            .cornerRadius(10)
                    }
                    .alert(isPresented: $showAgreementAlert) {
                        Alert(title: Text("提示"), message: Text("请先同意《用户协议》与《隐私协议》"), dismissButton: .default(Text("好的")))
                    }
                    .alert(isPresented: $showLoginErrorAlert) {
                        Alert(title: Text("登录失败"), message: Text("邮箱和密码不匹配"), dismissButton: .default(Text("重试")))
                    }
                }
                .padding(.horizontal)
                .padding(.bottom, 30)
                
                HStack {
                    Button(action: {
                        isAgreementChecked.toggle()
                    }) {
                        Image(systemName: isAgreementChecked ? "checkmark.square" : "square")
                            .foregroundColor(.pink)
                    }
                    Text("登录及代表同意《用户协议》与《隐私协议》")
                        .font(.footnote)
                        .foregroundColor(.gray)
                }
                .padding(.bottom, 20)
                
                Spacer()
                
                Text("其他登录方式")
                    .font(.footnote)
                    .foregroundColor(.gray)
                    .padding(.bottom, 10)
                
                HStack(spacing: 20) {
                    Image("wechat")
                        .resizable()
                        .frame(width: 50, height: 50)
                    
                    Image("apple")
                        .resizable()
                        .frame(width: 50, height: 50)
                    
                    Image("qq")
                        .resizable()
                        .frame(width: 50, height: 50)
                }
                .foregroundColor(Color.gray.opacity(0.7))
                .padding(.bottom, 30)
                
                // Navigation link to HomeView, triggered by login success
                NavigationLink(destination: HomeView().navigationBarBackButtonHidden(true), isActive: $isLoggedIn) {
                    EmptyView()
                }
            }
            .background(
                Image("login_back")
                    .resizable()
                    .scaledToFill()
                    .edgesIgnoringSafeArea(.all)
            )
        }
    }
    
    func login() {
        guard let url = URL(string: "https://69ee-2409-8931-b032-23-1942-57c8-de9c-79fd.ngrok-free.app/user/login") else { return }
        
        var request = URLRequest(url: url)
        request.httpMethod = "POST"
        request.addValue("application/json", forHTTPHeaderField: "Content-Type")
        
        let parameters: [String: Any] = [
            "email": email,
            "password": password
        ]
        
        request.httpBody = try? JSONSerialization.data(withJSONObject: parameters)
        
        let task = URLSession.shared.dataTask(with: request) { data, response, error in
            guard let data = data, error == nil else {
                print("错误:", error ?? "未知错误")
                return
            }
            
            if let httpResponse = response as? HTTPURLResponse {
                print("状态码:", httpResponse.statusCode)
                if httpResponse.statusCode == 200 {
                    // 登录成功，设置 isLoggedIn 为 true 跳转到 HomeView
                    DispatchQueue.main.async {
                        isLoggedIn = true
                    }
                } else {
                    // 显示错误提示
                    DispatchQueue.main.async {
                        showLoginErrorAlert = true
                    }
                }
            }
            
            let responseString = String(data: data, encoding: .utf8)
            print("响应内容:", responseString ?? "无响应")
        }
        
        task.resume()
    }
}



struct LoginView_Previews: PreviewProvider {
    static var previews: some View {
        LoginView()
    }
}
