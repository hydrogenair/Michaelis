import WidgetKit
import SwiftUI
struct Provider: TimelineProvider {
    func placeholder(in context: Context) -> SimpleEntry {
        SimpleEntry(date: Date())
    }

    func getSnapshot(in context: Context, completion: @escaping (SimpleEntry) -> Void) {
        let entry = SimpleEntry(date: Date())   
        completion(entry)
    }

    func getTimeline(in context: Context, completion: @escaping (Timeline<SimpleEntry>) -> Void) {
        let entries = [SimpleEntry(date: Date())]
        let timeline = Timeline(entries: entries, policy: .atEnd)
        completion(timeline)
    }
}

struct SimpleEntry: TimelineEntry {
    let date: Date
}

struct MyWidget: Widget {
    let kind: String = "lx.Michaelis.MyWidget"

    var body: some WidgetConfiguration {
        StaticConfiguration(kind: kind, provider: Provider()) { entry in
            WidgetEntryView(entry: entry)
                .widgetBackground(Color.white) // 这里替换背景色
        }
        .configurationDisplayName("My Widget")
        .description("This is a simple widget.")
        .supportedFamilies([.systemSmall, .systemMedium, .systemLarge])
    }
}

extension View {
    func widgetBackground(_ backgroundView: some View) -> some View {
        if #available(iOSApplicationExtension 17.0, *) {
            return containerBackground(for: .widget) {
                backgroundView
            }
        } else {
            return background(backgroundView)
        }
    }
}




struct TiredWidget: Widget {
    let kind: String = "lx.Michaelis.TiredWidget"

    var body: some WidgetConfiguration {
        StaticConfiguration(kind: kind, provider: Provider()) { entry in
            TiredWidgetEntryView(entry: entry)
                .widgetBackground2(Color.gray) // 使用 widgetBackground 并设置背景颜色
        }
        .configurationDisplayName("Tired Widget")
        .description("This is a tired widget.")
        .supportedFamilies([.systemSmall, .systemMedium, .systemLarge])
    }
}

extension View {
    func widgetBackground2(_ backgroundView: some View) -> some View {
        if #available(iOSApplicationExtension 17.0, *) {
            return containerBackground(for: .widget) {
                backgroundView
            }
        } else {
            return background(backgroundView)
        }
    }
}

