// Generated with prosr1 by compiler 0.9, template csharp.tmpl 0.1

using Microsoft.AspNetCore.SignalR;
using Microsoft.AspNetCore.SignalR.Client;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace Test.Package
{
	public interface ITestHubClient : ITestHubClientBase, IDisposable
	{
		Task InitializeHubConnection();
		Task CallFunctionOnHub(FunctionRequest message);
		Task CallTriggerOnHub();
		Task CallRefreshFunctionTreeOnHub();
	}

	public interface ITestHubClientBase
	{
		Task Function(FunctionResponse message);
		Task FunctionTreeRefreshed();
		Task FunctionCalled(FunctionRequest message);
	}

		
	public abstract class TestHubBase : Hub<ITestHubClientBase>
	{
		public async Task Function(FunctionRequest message)
		{
			var result = await HandleFunctionAsync(message).ConfigureAwait(false);

			await Clients
				.Caller
				.Function(result)
				.ConfigureAwait(false);
		}

		protected abstract Task<FunctionResponse> HandleFunctionAsync(FunctionRequest message);
		
		public async Task Trigger()
		{
			await HandleTriggerAsync().ConfigureAwait(false);
		}

		protected abstract Task HandleTriggerAsync();
		
		public async Task RefreshFunctionTree()
		{
			await HandleRefreshFunctionTreeAsync().ConfigureAwait(false);

			await Clients
				.Caller
				.FunctionTreeRefreshed()
				.ConfigureAwait(false);
		}

		protected abstract Task HandleRefreshFunctionTreeAsync();
		
		protected Task SendFunctionCalledOnAllAsync(FunctionRequest message)
		{
			return Clients.All.FunctionCalled(message);
		}

	}

	public class FunctionRequest
	{
		public string name { get; set; }
		public int page_number { get; set; }
		public int result_per_page { get; set; }
	}

	public class FunctionResponse
	{
		public IEnumerable<Test.Package2.Result> result { get; set; }
	}
}

namespace Test.Package2
{
	public class Result
	{
		public string url { get; set; }
		public string title { get; set; }
		public string snippets { get; set; }
	}
}