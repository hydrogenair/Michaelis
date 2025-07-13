import WidgetKit
import SwiftUI

@main
struct TiredWidget: Widget {
    let kind: String = "TiredWidget"

    var body: some WidgetConfiguration {
        StaticConfiguration(kind: kind, provider: Provider()) { entry in
            TiredWidgetEntryView(entry: entry)
        }
        .configurationDisplayName("疲惫提醒")
        .description("主人，今天看起来很疲惫呢。")
        .supportedFamilies([.systemSmall])
    }
}
