// Serilog.GoLogsinkSink::Serilog.GoLogsinkSink::NormalRenderedTextFormatter.cs
// 2021-02-25

namespace Serilog.GoLogsinkSink
{
	public class NormalRenderedTextFormatter : NormalTextFormatter
	{
		/// <summary>
		/// Initializes a new instance of the <see cref="NormalRenderedTextFormatter"/> class.
		/// </summary>
		public NormalRenderedTextFormatter()
		{
			IsRenderingMessage = true;
		}
	}
}
