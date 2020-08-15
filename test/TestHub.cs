// Generated with prosr1 by compiler 0.9, template csharp.tmpl 0.1

using Microsoft.AspNetCore.SignalR;
using Microsoft.AspNetCore.SignalR.Client;
using System;
using System.Threading.Tasks;


{
	public class FunctionRequest
	{
		public int result_per_page { get; set; }
		public int page_number { get; set; }
		public string name { get; set; }
	}

	public class FunctionResponse
	{
		public IEnumerable<Result> result { get; set; }
	}

	public class Result
	{
		public string snippets { get; set; }
		public string title { get; set; }
		public string url { get; set; }
	}
}