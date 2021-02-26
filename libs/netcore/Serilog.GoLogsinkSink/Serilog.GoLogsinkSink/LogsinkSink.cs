// Serilog.GoLogsinkSink::Serilog.GoLogsinkSink::LogsinkSink.cs
// 2021-02-25

namespace Serilog.GoLogsinkSink
{
	using System.IO;
	using Microsoft.Extensions.Logging;
	using System;
	using System.Collections.Generic;
	using System.Threading.Tasks;
	using Grpc.Net.Client;
	using Serilog.Events;
	using Serilog.Formatting;
	using Serilog.Sinks.PeriodicBatching;

	public class LogsinkSink : PeriodicBatchingSink
	{
		private readonly ITextFormatter _textFormatter;
		private readonly ILogger<LogsinkSink> _logger;
		private readonly LogTransfer.LogTransferClient _transferClient;
		private readonly GrpcChannel _channel;

		protected LogsinkSink(string requestUri, ILogger<LogsinkSink> logger, ITextFormatter textFormatter,
			IBatchedLogEventSink batchedSink, PeriodicBatchingSinkOptions options) : base(batchedSink,
			options)
		{
			_logger = logger;
			_textFormatter = textFormatter;

			_channel = GrpcChannel.ForAddress(requestUri);
			_transferClient = new LogTransfer.LogTransferClient(_channel);
		}

		public LogsinkSink(string requestUri, ILogger<LogsinkSink> logger, ITextFormatter textFormatter, int batchSizeLimit,
			TimeSpan period) : base(batchSizeLimit, period)
		{
			_logger = logger;
			_textFormatter = textFormatter;

			_channel = GrpcChannel.ForAddress(requestUri);
			_transferClient = new LogTransfer.LogTransferClient(_channel);
		}

		public LogsinkSink(string requestUri, ILogger<LogsinkSink> logger, ITextFormatter textFormatter, int batchSizeLimit,
			TimeSpan period, int queueLimit) : base(batchSizeLimit, period, queueLimit)
		{
			_logger = logger;
			_textFormatter = textFormatter;

			_channel = GrpcChannel.ForAddress(requestUri);
			_transferClient = new LogTransfer.LogTransferClient(_channel);
		}

		protected override async Task EmitBatchAsync(IEnumerable<LogEvent> events)
		{
			var client = _transferClient.SendLine();
			foreach (var logEvent in events)
			{
				var textWriter = new StringWriter();
				_textFormatter.Format(logEvent, textWriter);
				try
				{
					await client.RequestStream.WriteAsync(new LineMessage()
					{
						Line = textWriter.ToString(),
						Priority = 2
					});
				}
				catch (Exception ex)
				{
					Console.WriteLine($"got error: {ex}");
				}
			}

			await client.RequestStream.CompleteAsync();
		}

		protected override void Dispose(bool disposing)
		{
			_channel?.Dispose();
			base.Dispose(disposing);
		}
	}
}
