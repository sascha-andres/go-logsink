// Serilog.GoLogsinkSink::Serilog.GoLogsinkSink::a.cs
// 2021-02-25

namespace Serilog.GoLogsinkSink
{
	using System;
	using Microsoft.Extensions.Logging;
	using Serilog.Configuration;
	using Serilog.Events;
	using Serilog.Formatting;

	public static class LoggerSinkConfigurationExtensions
  {
    public static LoggerConfiguration AddLogsink(
      this LoggerSinkConfiguration sinkConfiguration,
      string requestUri,
      int batchPostingLimit = 1000,
      int? queueLimit = null,
      ILogger<LogsinkSink> logger = null,
      TimeSpan? period = null,
      ITextFormatter textFormatter = null,
      LogEventLevel restrictedToMinimumLevel = LevelAlias.Minimum)
    {
      if (sinkConfiguration == null) throw new ArgumentNullException(nameof(sinkConfiguration));

      // Default values
      period ??= TimeSpan.FromSeconds(2);
      textFormatter ??= new NormalRenderedTextFormatter();

      var sink = !queueLimit.HasValue
        ? new LogsinkSink(requestUri, logger, textFormatter, batchPostingLimit, period.Value)
        : new LogsinkSink(requestUri, logger, textFormatter, batchPostingLimit, period.Value, queueLimit.Value);

      return sinkConfiguration.Sink(sink, restrictedToMinimumLevel);
    }
  }
}
