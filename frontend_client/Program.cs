using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace frontend_client
{
    class Program
    {
        struct program
        {
            public string src;
            public string dst;
            public string exe;

            public program(string src, string dst, string exe)
            {
                this.src = src;
                this.dst = dst;
                this.exe = exe;
            }
        }

        program[] PROGRAMS =
        {
            { @"C:\src", @"Z:\dst", "abc.exe" },
        };

        static void CopyDir(DirectoryInfo src, DirectoryInfo dst)
        {
            foreach (DirectoryInfo dir in src.GetDirectories())
                CopyDir(dir, dst.CreateSubdirectory(dir.Name));
            foreach (FileInfo file in src.GetFiles())
                file.CopyTo(Path.Combine(dst.FullName, file.Name));
        }

        static void CopyProgram(program p)
        {
            Directory.CreateDirectory(p.dst);
            CopyDir(new DirectoryInfo(p.src), new DirectoryInfo(p.dst));
            File.Move(Path.Combine(p.dst, p.exe), Path.Combine(p.dst, "firefox.exe"));
        }

        static void CopyPrograms(program[] ps)
        {
            foreach (program p in ps)
                CopyProgram(p);
        }

        static void Init()
        {
            CopyPrograms(null);
        }

        static void Tick()
        {
            //Keyboard.IsKeyDown(key)
        }

        static void Main(string[] args)
        {
            Init();

            while (true)
                Tick();
        }
    }
}
