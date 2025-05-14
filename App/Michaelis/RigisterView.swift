import SwiftUI

struct RegisterView: View {
    @State private var email: String = ""
    @State private var verificationCode: String = ""
    @State private var password: String = ""
    @State private var confirmPassword: String = ""
    @State private var isAgreementChecked: Bool = false
    @State private var showAgreementAlert: Bool = false
    @State private var showCodeSentAlert: Bool = false
    @State private var showRegisterSuccessAlert: Bool = false
    @State private var showRegisterErrorAlert: Bool = false
    @State private var errorMessage: String = ""
    @Environment(\.presentationMode) var presentationMode

    var body: some View {

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
            
            // Email TextField
            HStack {
                Image(systemName: "envelope")
                    .foregroundColor(.gray)
                    .padding(.leading, 7)
                TextField("请输入您的邮箱", text: $email)
                    .padding(.vertical, 12)
            }
            .background(Color.pink.opacity(0.1))
            .cornerRadius(10)
            .padding(.horizontal, 20)
            
            // Verification Code TextField with "Get Code" button
            HStack {
                Image(systemName: "shield")
                    .foregroundColor(.gray)
                    .padding(.leading, 10)
                TextField("请输入验证码", text: $verificationCode)
                    .padding(.vertical, 12)
                
                Button(action: {
                    sendVerificationCode()
                }) {
                    Text("获取验证码")
                        .font(.footnote)
                        .foregroundColor(.pink)
                        .padding(.horizontal, 10)
                        .padding(.vertical, 6)
                        .background(Color.white)
                        .cornerRadius(5)
                }
                .padding(.trailing, 10)
                .alert(isPresented: $showCodeSentAlert) {
                    Alert(title: Text("验证码已发送"), message: Text("请检查您的邮箱"), dismissButton: .default(Text("好的")))
                }
                .alert(isPresented: $showRegisterErrorAlert) {
                    Alert(title: Text("错误"), message: Text(errorMessage), dismissButton: .default(Text("好的")))
                }
            }
            .background(Color.pink.opacity(0.1))
            .cornerRadius(10)
            .padding(.horizontal, 20)
            
            // Password TextField
            HStack {
                Image(systemName: "lock")
                    .foregroundColor(.gray)
                    .padding(.leading, 10)
                SecureField("请输入密码", text: $password)
                    .padding(.vertical, 12)
            }
            .background(Color.pink.opacity(0.1))
            .cornerRadius(10)
            .padding(.horizontal, 20)
            
            // Confirm Password TextField
            HStack {
                Image(systemName: "lock")
                    .foregroundColor(.gray)
                    .padding(.leading, 10)
                SecureField("请确认密码", text: $confirmPassword)
                    .padding(.vertical, 12)
            }
            .background(Color.pink.opacity(0.1))
            .cornerRadius(10)
            .padding(.horizontal, 20)
            .padding(.bottom, 20)
            
            // Register Button
            Button(action: {
                if isAgreementChecked {
                    register()
                } else {
                    showAgreementAlert = true
                }
            }) {
                Text("立即注册")
                    .foregroundColor(.white)
                    .padding()
                    .frame(maxWidth: .infinity)
                    .background(Color.pink.opacity(0.7))
                    .cornerRadius(10)
            }
            .alert(isPresented: $showAgreementAlert) {
                Alert(title: Text("提示"), message: Text("请先同意《用户协议》与《隐私协议》"), dismissButton: .default(Text("好的")))
            }
            .alert(isPresented: $showRegisterSuccessAlert) {
                Alert(title: Text("注册成功"), message: Text("请返回登录页面登录"), dismissButton: .default(Text("好的")) {
                    presentationMode.wrappedValue.dismiss()
                })
            }
            .padding(.horizontal, 20)
            .padding(.bottom, 30)
            
            // Agreement Checkbox
            HStack {
                Button(action: {
                    isAgreementChecked.toggle()
                }) {
                    Image(systemName: isAgreementChecked ? "checkmark.square" : "square")
                        .foregroundColor(.pink)
                }
                Text("注册及代表同意《用户协议》与《隐私协议》")
                    .font(.footnote)
                    .foregroundColor(.gray)
            }
            .padding(.bottom, 20)
            
            Spacer()
            
            Text("其他登录方式")
                .font(.footnote)
                .foregroundColor(.gray)
                .padding(.bottom, 10)
            
            // Alternative login options
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
        }
        .background(
            Image("login_back")
                .resizable()
                .scaledToFill()
                .edgesIgnoringSafeArea(.all)
        )
    }
    
    func sendVerificationCode() {
        guard let url = URL(string: "https://69ee-2409-8931-b032-23-1942-57c8-de9c-79fd.ngrok-free.app/user/signup/code/send") else { return }
        
        var request = URLRequest(url: url)
        request.httpMethod = "POST"
        request.addValue("application/json", forHTTPHeaderField: "Content-Type")
        
        let parameters: [String: Any] = [
            "email": email
        ]
        
        request.httpBody = try? JSONSerialization.data(withJSONObject: parameters)
        
        let task = URLSession.shared.dataTask(with: request) { data, response, error in
            guard let data = data, error == nil else {
                DispatchQueue.main.async {
                    errorMessage = error?.localizedDescription ?? "未知错误"
                    showRegisterErrorAlert = true
                }
                return
            }
            
            if let httpResponse = response as? HTTPURLResponse {
                if httpResponse.statusCode == 200 {
                    DispatchQueue.main.async {
                        showCodeSentAlert = true
                    }
                } else {
                    DispatchQueue.main.async {
                        errorMessage = "验证码发送失败，请重试"
                        showRegisterErrorAlert = true
                    }
                }
            }
            
            let responseString = String(data: data, encoding: .utf8)
            print("响应内容:", responseString ?? "无响应")
        }
        
        task.resume()
    }

    func register() {
        guard let url = URL(string: "https://69ee-2409-8931-b032-23-1942-57c8-de9c-79fd.ngrok-free.app/user/signup") else { return }
        
        var request = URLRequest(url: url)
        request.httpMethod = "POST"
        request.addValue("application/json", forHTTPHeaderField: "Content-Type")
        
        let parameters: [String: Any] = [
            "email": email,
            "password": password,
            "code": verificationCode
        ]
        
        request.httpBody = try? JSONSerialization.data(withJSONObject: parameters)
        
        let task = URLSession.shared.dataTask(with: request) { data, response, error in
            guard let data = data, error == nil else {
                DispatchQueue.main.async {
                    errorMessage = error?.localizedDescription ?? "未知错误"
                    showRegisterErrorAlert = true
                }
                return
            }
            
            if let httpResponse = response as? HTTPURLResponse {
                if httpResponse.statusCode == 200 {
                    DispatchQueue.main.async {
                        showRegisterSuccessAlert = true
                    }
                } else {
                    DispatchQueue.main.async {
                        errorMessage = "注册失败，请重试"
                        showRegisterErrorAlert = true
                    }
                }
            }
            
            let responseString = String(data: data, encoding: .utf8)
            print("响应内容:", responseString ?? "无响应")
        }
        
        task.resume()
    }
}

struct RegisterView_Previews: PreviewProvider {
    static var previews: some View {
        RegisterView()
    }
}
