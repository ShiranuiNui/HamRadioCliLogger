using System;

namespace HamRadioCliLogger
{
    public class QSO
    {
        public string MyCallSign { get; set; } = "JM8HBM";
        public string CallSign { get; set; }
        public long Time { get; set; } = DateTimeOffset.UtcNow.ToUnixTimeSeconds();
        public string Time_ISO8601 { get { return DateTimeOffset.FromUnixTimeSeconds(this.Time).ToLocalTime().ToString("yyyy-MM-ddTHH:mm:sszzzz"); } }
        public string Report { get; set; }
        public int Frequency { get; set; }
        public string Mode { get; set; }
        public bool IsRequestedQSLCard { get; set; }
        /*
        public bool IsSentQSLCard { get; private set; } = false;
        public bool IsReceivedQSLCard { get; private set; } = false;
        */
    }
}