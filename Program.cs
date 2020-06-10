using ConsoleAppFramework;
using ConsoleTables;
using Microsoft.Extensions.Hosting;
using System;
using System.Linq;
using System.Threading.Tasks;

namespace HamRadioCliLogger
{
    class Program : ConsoleAppBase
    {
        static async Task Main(string[] args)
        {
            await Host.CreateDefaultBuilder().RunConsoleAppFrameworkAsync<Program>(args);
        }
        [Command("new", "Create New QSO Logging")]
        public void RegistQSO(
            [Option("c", "name of send user.")] string callsign,
            [Option("r", "Sent Report like RST")] string report,
            [Option("f", "Frequency")] int freq,
            [Option("m", "Mode")] string mode,
            [Option("q", "Worked Station needs QSL Card?")] bool qsl = false)
        {
            var qso = new QSO() { CallSign = callsign, Report = report, Frequency = freq, Mode = mode, IsRequestedQSLCard = qsl };
            var rows = new QSO[] { qso };

            ConsoleTable
                .From<QSO>(rows)
                .Write(Format.Alternative);
        }
    }
}
