// Serilog.GoLogsinkSink::Serilog.GoLogsinkSink::NormalTextFormatter.cs
// 2021-02-25

using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Serilog.Debugging;
using Serilog.Events;
using Serilog.Formatting;
using Serilog.Formatting.Json;
using Serilog.Parsing;

namespace Serilog.GoLogsinkSink
{
	public class NormalTextFormatter : ITextFormatter
	{
		/// <summary>
		/// Gets or sets a value indicating whether the message is rendered into JSON.
		/// </summary>
		protected bool IsRenderingMessage { get; set; }

		/// <summary>
		/// Format the log event into the output.
		/// </summary>
		/// <param name="logEvent">The event to format.</param>
		/// <param name="output">The output.</param>
		public void Format(LogEvent logEvent, TextWriter output)
		{
			try
			{
				var buffer = new StringWriter();
				FormatContent(logEvent, buffer);

				// If formatting was successful, write to output
				output.WriteLine(buffer.ToString());
			}
			catch (Exception e)
			{
				LogNonFormattableEvent(logEvent, e);
			}
		}

		private void FormatContent(LogEvent logEvent, TextWriter output)
		{
			if (logEvent == null) throw new ArgumentNullException(nameof(logEvent));
			if (output == null) throw new ArgumentNullException(nameof(output));

			output.Write("{\"Timestamp\":\"");
			output.Write(logEvent.Timestamp.ToString("o"));

			output.Write("\",\"Level\":\"");
			output.Write(logEvent.Level);

			output.Write("\",\"MessageTemplate\":");
			JsonValueFormatter.WriteQuotedJsonString(logEvent.MessageTemplate.Text, output);

			if (IsRenderingMessage)
			{
				output.Write(",\"RenderedMessage\":");

				var message = logEvent.MessageTemplate.Render(logEvent.Properties);
				JsonValueFormatter.WriteQuotedJsonString(message, output);
			}

			if (logEvent.Exception != null)
			{
				output.Write(",\"Exception\":");
				JsonValueFormatter.WriteQuotedJsonString(logEvent.Exception.ToString(), output);
			}

			if (logEvent.Properties.Count != 0)
			{
				WriteProperties(logEvent.Properties, output);
			}

			// Better not to allocate an array in the 99.9% of cases where this is false
			var tokensWithFormat = logEvent.MessageTemplate.Tokens
				.OfType<PropertyToken>()
				.Where(pt => pt.Format != null);

			// ReSharper disable once PossibleMultipleEnumeration
			if (tokensWithFormat.Any())
			{
				// ReSharper disable once PossibleMultipleEnumeration
				WriteRenderings(tokensWithFormat.GroupBy(pt => pt.PropertyName), logEvent.Properties, output);
			}

			output.Write('}');
		}

		private static void WriteProperties(
			IReadOnlyDictionary<string, LogEventPropertyValue> properties,
			TextWriter output)
		{
			output.Write(",\"Properties\":{");

			var precedingDelimiter = "";

			foreach (var property in properties)
			{
				output.Write(precedingDelimiter);
				precedingDelimiter = ",";

				JsonValueFormatter.WriteQuotedJsonString(property.Key, output);
				output.Write(':');
				ValueFormatter.Instance.Format(property.Value, output);
			}

			output.Write('}');
		}

		private static void WriteRenderings(
			IEnumerable<IGrouping<string, PropertyToken>> tokensWithFormat,
			IReadOnlyDictionary<string, LogEventPropertyValue> properties,
			TextWriter output)
		{
			output.Write(",\"Renderings\":{");

			var rdelim = "";
			foreach (var ptoken in tokensWithFormat)
			{
				output.Write(rdelim);
				rdelim = ",";

				JsonValueFormatter.WriteQuotedJsonString(ptoken.Key, output);
				output.Write(":[");

				var fdelim = "";
				foreach (var format in ptoken)
				{
					output.Write(fdelim);
					fdelim = ",";

					output.Write("{\"Format\":");
					JsonValueFormatter.WriteQuotedJsonString(format.Format, output);

					output.Write(",\"Rendering\":");
					var sw = new StringWriter();
					format.Render(properties, sw);
					JsonValueFormatter.WriteQuotedJsonString(sw.ToString(), output);
					output.Write('}');
				}

				output.Write(']');
			}

			output.Write('}');
		}

		private static void LogNonFormattableEvent(LogEvent logEvent, Exception e)
		{
			SelfLog.WriteLine(
				"Event at {0} with message template {1} could not be formatted into JSON and will be dropped: {2}",
				logEvent.Timestamp.ToString("o"),
				logEvent.MessageTemplate.Text,
				e);
		}
	}
}
