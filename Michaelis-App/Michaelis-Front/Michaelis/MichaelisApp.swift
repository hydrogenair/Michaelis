

import SwiftUI

@main
struct MichaelisApp: App {
    @StateObject private var noteStore = NoteStore()

    var body: some Scene {
        WindowGroup {
            ContentView().environmentObject(NoteStore())
        }
    }
}
