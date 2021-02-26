// Serilog.GoLogsinkSink::Serilog.GoLogsinkSink::ValueFormatter.cs
// 2021-02-25

using Serilog.Formatting.Json;

namespace Serilog.GoLogsinkSink
{
	internal static class ValueFormatter
	{
		internal static readonly JsonValueFormatter Instance = new JsonValueFormatter();
	}
}
